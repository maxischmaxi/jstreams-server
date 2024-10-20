package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	lib "maxischmaxi/jstreams-server/lib"
	pb "maxischmaxi/jstreams-server/protos"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

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
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
