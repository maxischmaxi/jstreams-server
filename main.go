package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	lib "maxischmaxi/jstreams-server/lib"
	pb "maxischmaxi/jstreams-server/protos"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type assetsServer struct {
	pb.UnimplementedAssetsServiceServer
}

type championServer struct {
	pb.UnimplementedChampionsServiceServer
}

type accountServer struct {
	pb.UnimplementedAccountServiceServer
}

type masteriesServer struct {
	pb.UnimplementedMasteriesServiceServer
}

type matchesServer struct {
	pb.UnimplementedMatchesServiceServer
}

type summonerServer struct {
	pb.UnimplementedSummonerServiceServer
}

type entriesServer struct {
	pb.UnimplementedEntriesServiceServer
}

type tierServer struct {
	pb.UnimplementedTierServiceServer
}

type versionServer struct {
	pb.UnimplementedVersionServiceServer
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func (s *summonerServer) GetSummonerSpells(_ context.Context, in *pb.GetSummonerSpellsRequest) (*pb.GetSummonerSpellsResponse, error) {
	uri, err := lib.GetDDragonRoute(fmt.Sprintf("/cdn/%s/data/en_US/summoner.json", in.PatchVersion))

	if err != nil {
		log.Fatalf("failed to get summoner spells: %v", err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Fatalf("failed to get summoner spells: %v", err)
	}

	var spells pb.GetSummonerSpellsResponse
	err = json.NewDecoder(resp.Body).Decode(&spells)
	if err != nil {
		log.Fatalf("failed to decode summoner spells: %v", err)
	}

	return &spells, nil
}

func (s *summonerServer) GetSummonerByPuuidRequest(_ context.Context, in *pb.GetSummonerByPuuidRequest) (*pb.GetSummonerByPuuidResponse, error) {
	uri, err := lib.GetPlatformRoute(in.Region, fmt.Sprintf("/lol/summoner/v4/summoners/by-puuid/%s", in.Puuid))

	if err != nil {
		log.Fatalf("failed to get summoner by puuid: %v", err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Fatalf("failed to get summoner by puuid: %v", err)
	}

	var summoner pb.Summoner
	err = json.NewDecoder(resp.Body).Decode(&summoner)
	if err != nil {
		log.Fatalf("failed to decode summoner: %v", err)
	}

	return &pb.GetSummonerByPuuidResponse{
		Summoner: &summoner,
	}, nil
}

func (s *matchesServer) GetMatchIdsByPuuid(_ context.Context, in *pb.GetMatchIdsByPuuidRequest) (*pb.GetMatchIdsByPuuidResponse, error) {
	uri, err := lib.GetRegionalRoute(in.Region, fmt.Sprintf("/lol/match/v5/matches/by-puuid/%s/ids", in.Puuid))

	if err != nil {
		log.Fatalf("failed to get match ids by puuid: %v", err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Fatalf("failed to get match ids by puuid: %v", err)
	}

	var matchIds []string
	err = json.NewDecoder(resp.Body).Decode(&matchIds)
	if err != nil {
		log.Fatalf("failed to decode match ids: %v", err)
	}

	return &pb.GetMatchIdsByPuuidResponse{
		MatchIds: matchIds,
	}, nil
}

func (s *matchesServer) GetMatchTimeline(_ context.Context, in *pb.GetMatchTimelineRequest) (*pb.GetMatchTimelineResponse, error) {
	uri, err := lib.GetRegionalRoute(in.Region, fmt.Sprintf("/lol/match/v5/matches/%s/timeline", in.MatchId))

	if err != nil {
		log.Fatalf("failed to get match timeline: %v", err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Fatalf("failed to get match timeline: %v", err)
	}

	var timeline pb.Timeline
	err = json.NewDecoder(resp.Body).Decode(&timeline)
	if err != nil {
		log.Fatalf("failed to decode match timeline: %v", err)
	}

	return &pb.GetMatchTimelineResponse{
		Timeline: &timeline,
	}, nil
}

func (s *masteriesServer) GetChampionMateriesByPuuidByChampion(_ context.Context, in *pb.GetChampionMasteriesByChampionRequeset) (*pb.GetChampionMasteriesByChampionResponse, error) {
	uri, err := lib.GetPlatformRoute(in.Region, fmt.Sprintf("/lol/champion-mastery/v4/champion-masteries/by-puuid/%s/by-champion/%v", in.Puuid, in.ChampionId))

	if err != nil {
		log.Fatalf("failed to get champion masteries by puuid by champion: %v", err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Fatalf("failed to get champion masteries by puuid by champion: %v", err)
	}

	var mastery pb.ChampionMastery
	err = json.NewDecoder(resp.Body).Decode(&mastery)
	if err != nil {
		log.Fatalf("failed to decode champion mastery: %v", err)
	}

	return &pb.GetChampionMasteriesByChampionResponse{
		ChampionMastery: &mastery,
	}, nil
}

func (s *masteriesServer) GetChampionMasteriesByPuuid(_ context.Context, in *pb.GetChampionMasteriesRequeset) (*pb.GetChampionMasteriesResponse, error) {
	uri, err := lib.GetPlatformRoute(in.Region, fmt.Sprintf("/lol/champion-mastery/v4/champion-masteries/by-puuid/%s", in.Puuid))

	if err != nil {
		log.Fatalf("failed to get champion masteries by puuid: %v", err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Fatalf("failed to get champion masteries by puuid: %v", err)
	}

	var masteries []*pb.ChampionMastery

	err = json.NewDecoder(resp.Body).Decode(&masteries)
	if err != nil {
		log.Fatalf("failed to decode champion masteries: %v", err)
	}

	return &pb.GetChampionMasteriesResponse{
		ChampionMasteries: masteries,
	}, nil
}

func (s *championServer) GetChampions(_ context.Context, in *pb.GetChampionsRequest) (*pb.GetChampionsResponse, error) {
	uri, err := lib.GetDDragonRoute(fmt.Sprintf("/cdn/%s/data/en_US/champion.json", in.PatchVersion))

	if err != nil {
		log.Fatalf("failed to get champions: %v", err)
	}

	var response pb.GetChampionsResponse
	resp, err := http.Get(uri.String())
	if err != nil {
		log.Fatalf("failed to get champions: %v", err)
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatalf("failed to decode champions: %v", err)
	}

	return &response, nil
}

func (s *accountServer) GetAccountByGamenameAndTagline(_ context.Context, in *pb.GetAccountByGamenameAndTaglineRequest) (*pb.GetAccountByGamenameAndTaglineResponse, error) {
	uri, err := lib.GetRegionalRoute(in.Region, fmt.Sprintf("/riot/account/v1/accounts/by-riot-id/%s/%s", in.Gamename, in.Tagline))

	if err != nil {
		log.Fatalf("failed to get account by gamename and tagline: %v", err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Fatalf("failed to get account by gamename and tagline: %v", err)
	}

	var account pb.Account
	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		log.Fatalf("failed to decode account: %v", err)
	}

	return &pb.GetAccountByGamenameAndTaglineResponse{
		Account: &account,
	}, nil
}

func (s *accountServer) GetAccountByPuuid(_ context.Context, in *pb.GetAccountByPuuidRequest) (*pb.GetAccountByPuuidResponse, error) {
	uri, err := lib.GetRegionalRoute(in.Region, fmt.Sprintf("/riot/account/v1/accounts/by-puuid/%s", in.Puuid))

	if err != nil {
		log.Fatalf("failed to get account by puuid: %v", err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Fatalf("failed to get account by puuid: %v", err)
	}

	var account pb.Account
	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		log.Fatalf("failed to decode account: %v", err)
	}

	return &pb.GetAccountByPuuidResponse{
		Account: &account,
	}, nil
}

func (s *accountServer) GetAccountProfileIcon(_ context.Context, in *pb.GetAccountProfileIconRequest) (*pb.GetAccountProfileIconResponse, error) {
	var uri = fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img/profileicon/%v.png", in.PatchVersion, in.ProfileIconId)
	return &pb.GetAccountProfileIconResponse{
		Url: uri,
	}, nil
}

func (s *entriesServer) GetEntriesBySummoner(_ context.Context, in *pb.GetEntriesBySummonerRequest) (*pb.GetEntriesBySummonerResponse, error) {
	uri, err := lib.GetPlatformRoute(in.Region, fmt.Sprintf("/lol/league/v4/entries/by-summoner/%s", in.SummonerId))

	if err != nil {
		log.Fatalf("failed to get entries by summoner: %v", err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Fatalf("failed to get entries by summoner: %v", err)
	}

	var entries []*pb.SummonerEntry
	err = json.NewDecoder(resp.Body).Decode(&entries)
	if err != nil {
		log.Fatalf("failed to decode entries: %v", err)
	}

	return &pb.GetEntriesBySummonerResponse{
		Entries: entries,
	}, nil
}

func (s *tierServer) GetTierIcon(_ context.Context, in *pb.GetTierIconRequest) (*pb.GetTierIconResponse, error) {
	uri, err := lib.GetDDragonRoute(fmt.Sprintf("/cdn/%s/img/tier/%s.png", in.PatchVersion, in.Tier))

	if err != nil {
		log.Fatalf("failed to get tier icon: %v", err)
	}

	return &pb.GetTierIconResponse{
		Url: uri.String(),
	}, nil
}

func (s *versionServer) GetVersions(_ context.Context, in *pb.GetVersionsRequest) (*pb.GetVersionsResponse, error) {
	uri, err := lib.GetDDragonRoute("/api/versions.json")

	if err != nil {
		log.Fatalf("failed to get versions: %v", err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Fatalf("failed to get versions: %v", err)
	}

	var versions []string
	err = json.NewDecoder(resp.Body).Decode(&versions)

	if err != nil {
		log.Fatalf("failed to decode versions: %v", err)
	}

	return &pb.GetVersionsResponse{
		Versions: versions,
	}, nil
}

func (s *assetsServer) GetRuneIconRequest(_ context.Context, in *pb.GetRuneIconRequest) (*pb.GetRuneIconResponse, error) {
	uri, err := lib.GetDDragonRoute(fmt.Sprintf("/cdn/img/%s/img/%v", in.PatchVersion, in.Style))

	if err != nil {
		log.Fatalf("failed to get rune icon: %v", err)
	}

	return &pb.GetRuneIconResponse{
		Url: uri.String(),
	}, nil
}

func (s *assetsServer) GetSummonerSpellIcon(_ context.Context, in *pb.GetSummonerSpellIconRequest) (*pb.GetSummonerSpellIconResponse, error) {
	uri, err := lib.GetDDragonRoute(fmt.Sprintf("/cdn/%s/img/spell/%s", in.PatchVersion, in.Image.Full))

	if err != nil {
		log.Fatalf("failed to get summoner spell icon: %v", err)
	}

	return &pb.GetSummonerSpellIconResponse{
		Url: uri.String(),
	}, nil
}

func (s *assetsServer) GetItemAssetUrlRequest(_ context.Context, in *pb.GetItemAssetUrlRequest) (*pb.GetItemAssetUrlResponse, error) {
	uri, err := lib.GetDDragonRoute(fmt.Sprintf("/cdn/%s/data/en_US/%s.json", in.PatchVersion, in.ItemName))

	if err != nil {
		log.Fatalf("failed to get item asset url: %v", err)
	}

	return &pb.GetItemAssetUrlResponse{
		Url: uri.String(),
	}, nil
}

func (s *assetsServer) GetSpellAssetUrl(_ context.Context, in *pb.GetSpellAssetUrlRequest) (*pb.GetSpellAssetUrlResponse, error) {
	uri, err := lib.GetDDragonRoute(fmt.Sprintf("/cdn/%s/img/spell/%s.png", in.PatchVersion, in.SpellName))

	if err != nil {
		log.Fatalf("failed to get spell asset url: %v", err)
	}

	return &pb.GetSpellAssetUrlResponse{
		Url: uri.String(),
	}, nil
}

func (s *assetsServer) GetChampionAbilityAssetUrl(_ context.Context, in *pb.GetChampionAbilityAssetUrlRequest) (*pb.GetChampionAbilityAssetUrlResponse, error) {
	var championName = strings.ReplaceAll(in.ChampionName, " ", "")
	uri, err := lib.GetDDragonRoute(fmt.Sprintf("/cdn/%s/img/champion/%s_%v.png", in.PatchVersion, championName, in.AbilityNumber))

	if err != nil {
		log.Fatalf("failed to get champion ability asset url: %v", err)
	}

	return &pb.GetChampionAbilityAssetUrlResponse{
		Url: uri.String(),
	}, nil
}

func (s *assetsServer) GetChampionPassiveAssetUrl(_ context.Context, in *pb.GetChampionPassiveAssetUrlRequest) (*pb.GetChampionPassiveAssetUrlResponse, error) {
	var championName = strings.ReplaceAll(in.ChampionName, " ", "")
	uri, err := lib.GetDDragonRoute(fmt.Sprintf("/cdn/%s/img/passive/%s_P.png", in.PatchVersion, championName))

	if err != nil {
		log.Fatalf("failed to get champion passive asset url: %v", err)
	}

	return &pb.GetChampionPassiveAssetUrlResponse{
		Url: uri.String(),
	}, nil
}

func (s *assetsServer) GetChampionSquareAssetUrl(_ context.Context, in *pb.GetChampionSquareAssetUrlRequest) (*pb.GetChampionSquareAssetUrlResponse, error) {
	uri, err := lib.GetDDragonRoute(fmt.Sprintf("/cdn/%s/img/champion/%s", in.PatchVersion, in.ChampionName))

	if err != nil {
		log.Fatalf("failed to get champion square asset url: %v", err)
	}

	return &pb.GetChampionSquareAssetUrlResponse{
		Url: uri.String(),
	}, nil
}

func (s *assetsServer) GetChampionLoadingScreenAssetUrl(_ context.Context, in *pb.GetChampionLoadingScreenAssetUrlRequest) (*pb.GetChampionLoadingScreenAssetUrlResponse, error) {
	var championName = strings.ReplaceAll(in.ChampionName, " ", "")
	uri, err := lib.GetDDragonRoute(fmt.Sprintf("/cdn/img/champion/loading/%s_%v.jpg", championName, in.SkinNumber))

	if err != nil {
		log.Fatalf("failed to get champion loading screen asset url: %v", err)
	}

	return &pb.GetChampionLoadingScreenAssetUrlResponse{
		Url: uri.String(),
	}, nil
}

func (s *assetsServer) GetChampionSplashAssetUrl(_ context.Context, in *pb.GetChampionSplashAssetUrlRequest) (*pb.GetChampionSplashAssetUrlResponse, error) {
	var championName = strings.ReplaceAll(in.ChampionName, " ", "")
	uri, err := lib.GetDDragonRoute(fmt.Sprintf("/cdn/img/champion/splash/%s_%v.jpg", championName, in.SkinNumber))

	if err != nil {
		log.Fatalf("failed to get champion splash asset url: %v", err)
	}

	return &pb.GetChampionSplashAssetUrlResponse{
		Url: uri.String(),
	}, nil
}

func main() {
	flag.Parse()

	err := godotenv.Load(".env.development.local")
	if err != nil {
		log.Fatalf("failed to load .env.development.local: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSummonerServiceServer(s, &summonerServer{})
	pb.RegisterAccountServiceServer(s, &accountServer{})
	pb.RegisterMasteriesServiceServer(s, &masteriesServer{})
	pb.RegisterChampionsServiceServer(s, &championServer{})
	pb.RegisterMatchesServiceServer(s, &matchesServer{})
	pb.RegisterEntriesServiceServer(s, &entriesServer{})
	pb.RegisterTierServiceServer(s, &tierServer{})
	pb.RegisterVersionServiceServer(s, &versionServer{})
	pb.RegisterAssetsServiceServer(s, &assetsServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
