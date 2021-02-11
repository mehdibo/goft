package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"goft/pkg/ftapi"
	"strconv"
	"strings"
)

func getFilters(flags *pflag.FlagSet) (*map[string]string, error)  {
	filters := map[string]string{}
	// Map of flag => filter name, filter type
	flagFilters := map[string][2]string{
		"pool-year": {
			"pool_year",
			"StringSlice",
		},
		"pool-month": {
			"pool_month",
			"StringSlice",
		},
		"kind": {
			"kind",
			"StringSlice",
		},
		"primary-campus": {
			"primary_campus_id",
			"Int",
		},
	}
	for flagName, filter := range flagFilters {
		switch filter[1] {
		case "StringSlice":
			flagValue, err := flags.GetStringSlice(flagName)
			if err != nil {
				return nil, err
			}
			if len(flagValue) > 0 {
				filters[filter[0]] = strings.Join(flagValue, ",")
			}
		case "Int":
			flagValue, err := flags.GetInt(flagName)
			if err != nil {
				return nil, err
			}
			if flagValue > 0 {
				filters[filter[0]] = strconv.Itoa(flagValue)
			}
		}
	}
	if len(filters) > 0 {
		return &filters, nil
	}
	return nil, nil
}

func getSorts(queries []string) (*map[string]string, error) {
	sorts := map[string]string{}
	for _, query := range queries {
		tokens := strings.Split(query, ",")
		if len(tokens) != 2 {
			return nil, errors.New("invalid sort query: "+query)
		}
		tokens[1] = strings.ToLower(tokens[1])
		if tokens[1] != "asc" && tokens[1] != "desc" {
			return nil, errors.New("invalid sort order: "+tokens[1])
		}
		sorts[tokens[0]] = tokens[1]
	}
	if len(sorts) > 0 {
		return &sorts, nil
	}
	return nil, nil
}

func filterNils(users []*ftapi.User) []*ftapi.User {
	for i := 0; i < len(users); {
		if users[i] != nil {
			i++
			continue
		}
		if i < len(users) - 1 {
			copy(users[i:], users[i+1:])
		}
		users[len(users)-1] = nil
		users = users[:len(users)-1]

	}
	return users
}

// NewListUsersCmd create new list users cmd
func NewListUsersCmd(api *ftapi.APIInterface) *cobra.Command {
	cmd := cobra.Command{
		Use:   "ls",
		Short: "List users",
		Long: `Use this command to list users.
You can use and combine the flags to filter through users.

Available fields for sort:
id, login, email, created_at, updated_at, first_name
last_name, pool_year, pool_month, kind, slack_login
last_seen_at, password_changed_at
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			showAnonym, err := cmd.Flags().GetBool("show-anonym")
			if err != nil {
				return err
			}
			filters, err := getFilters(cmd.Flags())
			if err != nil {
				return err
			}
			sortFlag, err := cmd.Flags().GetStringArray("sort")
			if err != nil {
				return err
			}
			sort, err := getSorts(sortFlag)
			if err != nil {
				return err
			}
			// Load all users first
			page := 1
			var users []*ftapi.User
			for {
				tmpUsers, err := (*api).GetUsers(page, filters, sort)
				if err != nil {
					return err
				}
				// Response is empty, means last page
				if len(tmpUsers) == 0 {
					break
				}
				users = append(users, tmpUsers...)
				page++
			}
			if !showAnonym {
				for i, user := range users {
					if user.IsAnonymized(){
						users[i] = nil
					}
				}
				users = filterNils(users)
			}
			for _, user := range users {
				fmt.Println(user)
			}
			return nil
		},
	}
	cmd.Flags().StringSlice("pool-year", nil, "Filter by pool year (values separated by a comma)")
	cmd.Flags().StringSlice("pool-month", nil, "Filter by pool month (values separated by a comma)")
	cmd.Flags().StringSlice("kind", nil, "Filter by kind, options are: student, admin and external (values separated by a comma)")
	cmd.Flags().Int("primary-campus", 0, "Filter by primary campus id")
	cmd.Flags().StringArray("sort", nil, "Sort the users, format: field_name,asc|desc")
	cmd.Flags().Bool("show-anonym", false, "Use to show anonymized users")
	return &cmd
}
var listUsersCmd = NewListUsersCmd(&API)

func init() {
	usersCmd.AddCommand(listUsersCmd)
}
