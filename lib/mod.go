package lib

import (
	"fmt"
	pb "maxischmaxi/jstreams-server/protos"
	"net/url"
	"os"
)

func GetPlatformRoute(region pb.PlatformRoutingValues, route string) (url.URL, error) {
	var uri = fmt.Sprintf("https://%s.api.riotgames.com%s", region, route)
	u, err := url.Parse(uri)

	if err != nil {
		return url.URL{}, err
	}

	q := u.Query()
	q.Set("api_key", os.Getenv("RIOT_API_KEY"))
	u.RawQuery = q.Encode()

	return *u, nil
}

func GetRegionalRoute(region pb.RegionalRoutingValues, route string) (url.URL, error) {
	var uri = fmt.Sprintf("https://%s.api.riotgames.com%s", region, route)
	u, err := url.Parse(uri)

	if err != nil {
		return url.URL{}, err
	}

	q := u.Query()
	q.Set("api_key", os.Getenv("RIOT_API_KEY"))
	u.RawQuery = q.Encode()

	return *u, nil
}

func GetDDragonRoute(route string) (url.URL, error) {
	var uri = fmt.Sprintf("https://ddragon.leagueoflegends.com%s", route)
	u, err := url.Parse(uri)

	if err != nil {
		return url.URL{}, err
	}

	return *u, nil
}
