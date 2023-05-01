package abios

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/metsawyr/abios-api/internal/config"
)

type AbiosClient struct {
	config     *config.Config
	httpClient *http.Client
}

type AbiosSeries struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Participants []struct {
		Roster struct {
			Id int `json:"id"`
		} `json:"roster"`
	} `json:"participants"`
}

type AbiosRoster struct {
	Id   int `json:"id"`
	Team struct {
		Id int `json:"id"`
	} `json:"team"`
	LineUp struct {
		Players []struct {
			Id int `json:"id"`
		} `json:"players"`
	} `json:"line_up"`
}

type AbiosPlayer struct {
	Id       int    `json:"id"`
	Nickname string `json:"nick_name"`
}

type AbiosTeam struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"abbreviation"`
}

func NewAbiosClient(config *config.Config) AbiosClient {
	return AbiosClient{
		config:     config,
		httpClient: &http.Client{},
	}
}

func (c AbiosClient) LiveSeries(ctx context.Context) (*[]AbiosSeries, error) {
	return get[[]AbiosSeries](c.config, c.httpClient, ctx, "series", "lifecycle=live")
}

func (c AbiosClient) LiveRosters(ctx context.Context) (*[]AbiosRoster, error) {
	series, err := c.LiveSeries(ctx)
	if err != nil {
		return nil, err
	}

	rosterIds := make([]string, 0)
	for _, currentSeries := range *series {
		for _, participant := range currentSeries.Participants {
			rosterIds = append(rosterIds, fmt.Sprint(participant.Roster.Id))
		}
	}

	return get[[]AbiosRoster](c.config, c.httpClient, ctx, "rosters", fmt.Sprintf("id<={%v}", strings.Join(rosterIds, ",")))
}

func (c AbiosClient) LivePlayers(ctx context.Context) (*[]AbiosPlayer, error) {
	rosters, err := c.LiveRosters(ctx)
	if err != nil {
		return nil, err
	}

	playerIds := make([]string, 0)
	for _, roster := range *rosters {
		for _, player := range roster.LineUp.Players {
			playerIds = append(playerIds, fmt.Sprint(player.Id))
		}
	}

	return get[[]AbiosPlayer](c.config, c.httpClient, ctx, "players", fmt.Sprintf("id<={%v}", strings.Join(playerIds, ",")))
}

func (c AbiosClient) LiveTeams(ctx context.Context) (*[]AbiosTeam, error) {
	rosters, err := c.LiveRosters(ctx)
	if err != nil {
		return nil, err
	}

	teamIds := make([]string, 0)
	for _, roster := range *rosters {
		teamIds = append(teamIds, fmt.Sprint(roster.Team.Id))
	}

	return get[[]AbiosTeam](c.config, c.httpClient, ctx, "teams", fmt.Sprintf("id<={%v}", strings.Join(teamIds, ",")))
}

func get[T interface{}](
	config *config.Config,
	httpClient *http.Client,
	ctx context.Context,
	path string,
	filter string,
) (*T, error) {
	uri, err := url.JoinPath(config.Abios.ApiUri, path)
	if err != nil {
		return nil, err
	}
	log.Println("URI: ", uri)

	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	request = request.WithContext(ctx)
	request.Header.Set("Abios-Secret", config.Abios.Secret)
	if len(filter) > 0 {
		query := request.URL.Query()
		query.Add("filter", filter)
		request.URL.RawQuery = query.Encode()
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
