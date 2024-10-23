package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	accountv1 "maxischmaxi/jstreams-server/gen/account/v1"
	"maxischmaxi/jstreams-server/gen/account/v1/accountv1connect"
	assetsv1 "maxischmaxi/jstreams-server/gen/assets/v1"
	"maxischmaxi/jstreams-server/gen/assets/v1/assetsv1connect"
	championsv1 "maxischmaxi/jstreams-server/gen/champions/v1"
	"maxischmaxi/jstreams-server/gen/champions/v1/championsv1connect"
	entriesv1 "maxischmaxi/jstreams-server/gen/entries/v1"
	"maxischmaxi/jstreams-server/gen/entries/v1/entriesv1connect"
	masteriesv1 "maxischmaxi/jstreams-server/gen/masteries/v1"
	"maxischmaxi/jstreams-server/gen/masteries/v1/masteriesv1connect"
	matchesv1 "maxischmaxi/jstreams-server/gen/matches/v1"
	"maxischmaxi/jstreams-server/gen/matches/v1/matchesv1connect"
	summonerv1 "maxischmaxi/jstreams-server/gen/summoner/v1"
	"maxischmaxi/jstreams-server/gen/summoner/v1/summonerv1connect"
	tierv1 "maxischmaxi/jstreams-server/gen/tier/v1"
	"maxischmaxi/jstreams-server/gen/tier/v1/tierv1connect"
	versionv1 "maxischmaxi/jstreams-server/gen/version/v1"
	"maxischmaxi/jstreams-server/gen/version/v1/versionv1connect"

	"connectrpc.com/connect"

	"github.com/joho/godotenv"
)

type assetsServer struct{}

type championsServer struct{}

type accountServer struct{}

type masteriesServer struct{}

type matchesServer struct{}

type summonerServer struct{}

type entriesServer struct{}

type tierServer struct{}

type versionServer struct{}

func (s *summonerServer) GetSummonerSpells(
	_ context.Context,
	in *connect.Request[summonerv1.GetSummonerSpellsRequest],
) (*connect.Response[summonerv1.GetSummonerSpellsResponse], error) {
	uri, err := GetDDragonRoute(fmt.Sprintf("/cdn/%s/data/en_US/summoner.json", in.Msg.PatchVersion))

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var spells summonerv1.GetSummonerSpellsResponse
	err = json.NewDecoder(resp.Body).Decode(&spells)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&spells), nil
}

func (s *summonerServer) GetSummonerByPuuid(
	_ context.Context,
	in *connect.Request[summonerv1.GetSummonerByPuuidRequest],
) (*connect.Response[summonerv1.GetSummonerByPuuidResponse], error) {
	path := fmt.Sprintf("/lol/summoner/v4/summoners/by-puuid/%s", in.Msg.Puuid)
	uri, err := GetPlatformRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var summoner summonerv1.Summoner
	err = json.NewDecoder(resp.Body).Decode(&summoner)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&summonerv1.GetSummonerByPuuidResponse{
		Summoner: &summoner,
	}), nil
}

func (s *matchesServer) GetMatchIdsByPuuid(
	_ context.Context,
	in *connect.Request[matchesv1.GetMatchIdsByPuuidRequest],
) (*connect.Response[matchesv1.GetMatchIdsByPuuidResponse], error) {
	path := fmt.Sprintf("/lol/match/v5/matches/by-puuid/%s/ids", in.Msg.Puuid)
	uri, err := GetRegionalRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var matchIds []string
	err = json.NewDecoder(resp.Body).Decode(&matchIds)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&matchesv1.GetMatchIdsByPuuidResponse{
		MatchIds: matchIds,
	}), nil
}

func (s *matchesServer) GetMatchByMatchId(
	_ context.Context,
	in *connect.Request[matchesv1.GetMatchByMatchIdRequest],
) (*connect.Response[matchesv1.GetMatchByMatchIdResponse], error) {
	return connect.NewResponse(&matchesv1.GetMatchByMatchIdResponse{}), nil
}

func (s *matchesServer) GetMatchTimeline(
	_ context.Context,
	in *connect.Request[matchesv1.GetMatchTimelineRequest],
) (*connect.Response[matchesv1.GetMatchTimelineResponse], error) {
	path := fmt.Sprintf("/lol/match/v5/matches/%s/timeline", in.Msg.MatchId)
	uri, err := GetRegionalRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var timeline matchesv1.Timeline
	err = json.NewDecoder(resp.Body).Decode(&timeline)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&matchesv1.GetMatchTimelineResponse{
		Timeline: &timeline,
	}), nil
}

func (s *masteriesServer) GetChampionMasteriesByPuuidByChampion(
	_ context.Context,
	in *connect.Request[masteriesv1.GetChampionMasteriesByChampionRequeset],
) (*connect.Response[masteriesv1.GetChampionMasteriesByChampionResponse], error) {
	return connect.NewResponse(&masteriesv1.GetChampionMasteriesByChampionResponse{}), nil
}

func (s *masteriesServer) GetChampionMateriesByPuuidByChampion(
	_ context.Context,
	in *connect.Request[masteriesv1.GetChampionMasteriesByChampionRequeset],
) (*connect.Response[masteriesv1.GetChampionMasteriesByChampionResponse], error) {
	path := fmt.Sprintf("/lol/champion-mastery/v4/champion-masteries/by-puuid/%s/by-champion/%v", in.Msg.Puuid, in.Msg.ChampionId)
	uri, err := GetPlatformRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var mastery masteriesv1.ChampionMastery
	err = json.NewDecoder(resp.Body).Decode(&mastery)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&masteriesv1.GetChampionMasteriesByChampionResponse{
		ChampionMastery: &mastery,
	}), nil
}

func (s *masteriesServer) GetChampionMasteriesByPuuid(
	_ context.Context,
	in *connect.Request[masteriesv1.GetChampionMasteriesRequeset],
) (*connect.Response[masteriesv1.GetChampionMasteriesResponse], error) {
	path := fmt.Sprintf("/lol/champion-mastery/v4/champion-masteries/by-puuid/%s", in.Msg.Puuid)
	uri, err := GetPlatformRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var masteries []*masteriesv1.ChampionMastery

	err = json.NewDecoder(resp.Body).Decode(&masteries)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&masteriesv1.GetChampionMasteriesResponse{
		ChampionMasteries: masteries,
	}), nil
}

func (s *championsServer) GetChampions(
	_ context.Context,
	in *connect.Request[championsv1.GetChampionsRequest],
) (*connect.Response[championsv1.GetChampionsResponse], error) {
	path := fmt.Sprintf("/cdn/%s/data/en_US/champion.json", in.Msg.PatchVersion)
	uri, err := GetDDragonRoute(path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var response championsv1.GetChampionsResponse
	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&response), nil
}

func (s *accountServer) GetAccountByGamenameAndTagline(
	_ context.Context,
	in *connect.Request[accountv1.GetAccountByGamenameAndTaglineRequest],
) (*connect.Response[accountv1.GetAccountByGamenameAndTaglineResponse], error) {
	path := fmt.Sprintf("/riot/account/v1/accounts/by-riot-id/%s/%s", in.Msg.Gamename, in.Msg.Tagline)
	uri, err := GetRegionalRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var account accountv1.Account
	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&accountv1.GetAccountByGamenameAndTaglineResponse{
		Account: &account,
	}), nil
}

func (s *accountServer) GetAccountByPuuid(
	_ context.Context,
	in *connect.Request[accountv1.GetAccountByPuuidRequest],
) (*connect.Response[accountv1.GetAccountByPuuidResponse], error) {
	path := fmt.Sprintf("/riot/account/v1/accounts/by-puuid/%s", in.Msg.Puuid)
	uri, err := GetRegionalRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var account accountv1.Account
	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&accountv1.GetAccountByPuuidResponse{
		Account: &account,
	}), nil
}

func (s *accountServer) GetAccountProfileIcon(
	_ context.Context,
	in *connect.Request[accountv1.GetAccountProfileIconRequest],
) (*connect.Response[accountv1.GetAccountProfileIconResponse], error) {
	var uri = fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img/profileicon/%v.png", in.Msg.PatchVersion, in.Msg.ProfileIconId)
	return connect.NewResponse(&accountv1.GetAccountProfileIconResponse{
		Url: uri,
	}), nil
}

func (s *entriesServer) GetEntriesBySummoner(
	_ context.Context,
	in *connect.Request[entriesv1.GetEntriesBySummonerRequest],
) (*connect.Response[entriesv1.GetEntriesBySummonerResponse], error) {
	path := fmt.Sprintf("/lol/league/v4/entries/by-summoner/%s", in.Msg.SummonerId)
	uri, err := GetPlatformRoute(in.Msg.Region, path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var entries []*entriesv1.SummonerEntry
	err = json.NewDecoder(resp.Body).Decode(&entries)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&entriesv1.GetEntriesBySummonerResponse{
		Entries: entries,
	}), nil
}

func (s *tierServer) GetTierIcon(
	_ context.Context,
	in *connect.Request[tierv1.GetTierIconRequest],
) (*connect.Response[tierv1.GetTierIconResponse], error) {
	path := fmt.Sprintf("/cdn/%s/img/tier/%s.png", in.Msg.PatchVersion, in.Msg.Tier)
	uri, err := GetDDragonRoute(path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&tierv1.GetTierIconResponse{
		Url: uri.String(),
	}), nil
}

func (s *versionServer) GetVersions(
	_ context.Context,
	in *connect.Request[versionv1.GetVersionsRequest],
) (*connect.Response[versionv1.GetVersionsResponse], error) {
	uri, err := GetDDragonRoute("/api/versions.json")

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	var versions []string
	err = json.NewDecoder(resp.Body).Decode(&versions)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&versionv1.GetVersionsResponse{
		Versions: versions,
	}), nil
}

func (s *assetsServer) GetRuneIcon(
	_ context.Context,
	in *connect.Request[assetsv1.GetRuneIconRequest],
) (*connect.Response[assetsv1.GetRuneIconResponse], error) {
	return connect.NewResponse(&assetsv1.GetRuneIconResponse{
		Url: "",
	}), nil
}

func (s *assetsServer) GetItemAssetUrl(
	_ context.Context,
	in *connect.Request[assetsv1.GetItemAssetUrlRequest],
) (*connect.Response[assetsv1.GetItemAssetUrlResponse], error) {
	return connect.NewResponse(&assetsv1.GetItemAssetUrlResponse{
		Url: "",
	}), nil
}

func (s *assetsServer) GetRuneIconRequest(
	_ context.Context,
	in *connect.Request[assetsv1.GetRuneIconRequest],
) (*connect.Response[assetsv1.GetRuneIconResponse], error) {
	path := fmt.Sprintf("/cdn/img/%s/img/%v", in.Msg.PatchVersion, in.Msg.Style)
	uri, err := GetDDragonRoute(path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetRuneIconResponse{
		Url: uri.String(),
	}), nil
}

func (s *assetsServer) GetSummonerSpellIcon(
	_ context.Context,
	in *connect.Request[assetsv1.GetSummonerSpellIconRequest],
) (*connect.Response[assetsv1.GetSummonerSpellIconResponse], error) {
	path := fmt.Sprintf("/cdn/%s/img/spell/%s", in.Msg.PatchVersion, in.Msg.Image.Full)
	uri, err := GetDDragonRoute(path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetSummonerSpellIconResponse{
		Url: uri.String(),
	}), nil
}

func (s *assetsServer) GetItemAssetUrlRequest(
	_ context.Context,
	in *connect.Request[assetsv1.GetItemAssetUrlRequest],
) (*connect.Response[assetsv1.GetItemAssetUrlResponse], error) {
	path := fmt.Sprintf("/cdn/%s/data/en_US/%s.json", in.Msg.PatchVersion, in.Msg.ItemName)
	uri, err := GetDDragonRoute(path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetItemAssetUrlResponse{
		Url: uri.String(),
	}), nil
}

func (s *assetsServer) GetSpellAssetUrl(
	_ context.Context,
	in *connect.Request[assetsv1.GetSpellAssetUrlRequest],
) (*connect.Response[assetsv1.GetSpellAssetUrlResponse], error) {
	uri, err := GetDDragonRoute(fmt.Sprintf("/cdn/%s/img/spell/%s.png", in.Msg.PatchVersion, in.Msg.SpellName))

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetSpellAssetUrlResponse{
		Url: uri.String(),
	}), nil
}

func (s *assetsServer) GetChampionAbilityAssetUrl(
	_ context.Context,
	in *connect.Request[assetsv1.GetChampionAbilityAssetUrlRequest],
) (*connect.Response[assetsv1.GetChampionAbilityAssetUrlResponse], error) {
	var championName = strings.ReplaceAll(in.Msg.ChampionName, " ", "")
	uri, err := GetDDragonRoute(fmt.Sprintf("/cdn/%s/img/champion/%s_%v.png", in.Msg.PatchVersion, championName, in.Msg.AbilityNumber))

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetChampionAbilityAssetUrlResponse{
		Url: uri.String(),
	}), nil
}

func (s *assetsServer) GetChampionPassiveAssetUrl(
	_ context.Context,
	in *connect.Request[assetsv1.GetChampionPassiveAssetUrlRequest],
) (*connect.Response[assetsv1.GetChampionPassiveAssetUrlResponse], error) {
	var championName = strings.ReplaceAll(in.Msg.ChampionName, " ", "")
	path := fmt.Sprintf("/cdn/%s/img/passive/%s_P.png", in.Msg.PatchVersion, championName)
	uri, err := GetDDragonRoute(path)

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetChampionPassiveAssetUrlResponse{
		Url: uri.String(),
	}), nil
}

func (s *assetsServer) GetChampionSquareAssetUrl(
	_ context.Context,
	in *connect.Request[assetsv1.GetChampionSquareAssetUrlRequest],
) (*connect.Response[assetsv1.GetChampionSquareAssetUrlResponse], error) {
	uri, err := GetDDragonRoute(fmt.Sprintf("/cdn/%s/img/champion/%s", in.Msg.PatchVersion, in.Msg.ChampionName))

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetChampionSquareAssetUrlResponse{
		Url: uri.String(),
	}), nil
}

func (s *assetsServer) GetChampionLoadingScreenAssetUrl(
	_ context.Context,
	in *connect.Request[assetsv1.GetChampionLoadingScreenAssetUrlRequest],
) (*connect.Response[assetsv1.GetChampionLoadingScreenAssetUrlResponse], error) {
	var championName = strings.ReplaceAll(in.Msg.ChampionName, " ", "")
	uri, err := GetDDragonRoute(fmt.Sprintf("/cdn/img/champion/loading/%s_%v.jpg", championName, in.Msg.SkinNumber))

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetChampionLoadingScreenAssetUrlResponse{
		Url: uri.String(),
	}), nil
}

func (s *assetsServer) GetChampionSplashAssetUrl(
	_ context.Context,
	in *connect.Request[assetsv1.GetChampionSplashAssetUrlRequest],
) (*connect.Response[assetsv1.GetChampionSplashAssetUrlResponse], error) {
	var championName = strings.ReplaceAll(in.Msg.ChampionName, " ", "")
	uri, err := GetDDragonRoute(fmt.Sprintf("/cdn/img/champion/splash/%s_%v.jpg", championName, in.Msg.SkinNumber))

	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&assetsv1.GetChampionSplashAssetUrlResponse{
		Url: uri.String(),
	}), nil
}

func main() {
	err := godotenv.Load(".env.development.local")
	if err != nil {
		log.Fatalf("failed to load .env.development.local: %v", err)
	}

	mux := http.NewServeMux()
	interceptors := connect.WithInterceptors(LoggingInterceptor())

	champions := &championsServer{}
	path, handler := championsv1connect.NewChampionsServiceHandler(champions, interceptors)
	mux.Handle(path, handler)

	summoner := &summonerServer{}
	path, handler = summonerv1connect.NewSummonerServiceHandler(summoner, interceptors)
	mux.Handle(path, handler)

	matches := &matchesServer{}
	path, handler = matchesv1connect.NewMatchesServiceHandler(matches, interceptors)
	mux.Handle(path, handler)

	entries := &entriesServer{}
	path, handler = entriesv1connect.NewEntriesServiceHandler(entries, interceptors)
	mux.Handle(path, handler)

	assets := &assetsServer{}
	path, handler = assetsv1connect.NewAssetsServiceHandler(assets, interceptors)
	mux.Handle(path, handler)

	tier := &tierServer{}
	path, handler = tierv1connect.NewTierServiceHandler(tier, interceptors)
	mux.Handle(path, handler)

	account := &accountServer{}
	path, handler = accountv1connect.NewAccountServiceHandler(account, interceptors)
	mux.Handle(path, handler)

	version := &versionServer{}
	path, handler = versionv1connect.NewVersionServiceHandler(version, interceptors)
	mux.Handle(path, handler)

	masteris := &masteriesServer{}
	path, handler = masteriesv1connect.NewMasteriesServiceHandler(masteris, interceptors)
	mux.Handle(path, handler)

	log.Fatal(http.ListenAndServe(
		"localhost:8080",
		mux,
	))
}

func GetPlatformRoute(region summonerv1.PlatformRoutingValues, route string) (url.URL, error) {
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

func GetRegionalRoute(region accountv1.RegionalRoutingValues, route string) (url.URL, error) {
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

func LoggingInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			log.Printf("[%s] %s %s %s", req.HTTPMethod(), req.Peer().Protocol, req.Peer().Addr, req.Peer().Query)
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
