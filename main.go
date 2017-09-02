package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	Provider struct {
		Name      string `json:"name"`
		Appid     int    `json:"appid"`
		Version   int    `json:"version"`
		Steamid   string `json:"steamid"`
		Timestamp int    `json:"timestamp"`
	} `json:"provider"`
	Player struct {
		Steamid  string `json:"steamid"`
		Name     string `json:"name"`
		Activity string `json:"activity"`
	} `json:"player"`
	Previously struct {
		Map    bool `json:"map"`
		Round  bool `json:"round"`
		Player struct {
			Clan         string `json:"clan"`
			ObserverSlot int    `json:"observer_slot"`
			Team         string `json:"team"`
			Activity     string `json:"activity"`
			State        struct {
				Health      int  `json:"health"`
				Armor       int  `json:"armor"`
				Helmet      bool `json:"helmet"`
				Flashed     int  `json:"flashed"`
				Smoked      int  `json:"smoked"`
				Burning     int  `json:"burning"`
				Money       int  `json:"money"`
				RoundKills  int  `json:"round_kills"`
				RoundKillhs int  `json:"round_killhs"`
				EquipValue  int  `json:"equip_value"`
			} `json:"state"`
			Weapons struct {
				Weapon0 struct {
					Name     string `json:"name"`
					Paintkit string `json:"paintkit"`
					Type     string `json:"type"`
					State    string `json:"state"`
				} `json:"weapon_0"`
				Weapon1 struct {
					Name        string `json:"name"`
					Paintkit    string `json:"paintkit"`
					Type        string `json:"type"`
					AmmoClip    int    `json:"ammo_clip"`
					AmmoClipMax int    `json:"ammo_clip_max"`
					AmmoReserve int    `json:"ammo_reserve"`
					State       string `json:"state"`
				} `json:"weapon_1"`
				Weapon2 struct {
					Name     string `json:"name"`
					Paintkit string `json:"paintkit"`
					Type     string `json:"type"`
					State    string `json:"state"`
				} `json:"weapon_2"`
				Weapon3 struct {
					Name        string `json:"name"`
					Paintkit    string `json:"paintkit"`
					Type        string `json:"type"`
					AmmoClip    int    `json:"ammo_clip"`
					AmmoClipMax int    `json:"ammo_clip_max"`
					AmmoReserve int    `json:"ammo_reserve"`
					State       string `json:"state"`
				} `json:"weapon_3"`
			} `json:"weapons"`
			MatchStats struct {
				Kills   int `json:"kills"`
				Assists int `json:"assists"`
				Deaths  int `json:"deaths"`
				Mvps    int `json:"mvps"`
				Score   int `json:"score"`
			} `json:"match_stats"`
		} `json:"player"`
	} `json:"previously"`
	Auth struct {
		Token string `json:"token"`
	} `json:"auth"`
}

type AutoGenerated struct {
	Provider struct {
		Name      string `json:"name"`
		Appid     int    `json:"appid"`
		Version   int    `json:"version"`
		Steamid   string `json:"steamid"`
		Timestamp int    `json:"timestamp"`
	} `json:"provider"`
	Map struct {
		Mode   string `json:"mode"`
		Name   string `json:"name"`
		Phase  string `json:"phase"`
		Round  int    `json:"round"`
		TeamCt struct {
			Score                int `json:"score"`
			TimeoutsRemaining    int `json:"timeouts_remaining"`
			MatchesWonThisSeries int `json:"matches_won_this_series"`
		} `json:"team_ct"`
		TeamT struct {
			Score                int `json:"score"`
			TimeoutsRemaining    int `json:"timeouts_remaining"`
			MatchesWonThisSeries int `json:"matches_won_this_series"`
		} `json:"team_t"`
		NumMatchesToWinSeries int `json:"num_matches_to_win_series"`
		CurrentSpectators     int `json:"current_spectators"`
		SouvenirsTotal        int `json:"souvenirs_total"`
	} `json:"map"`
	Round struct {
		Phase   string `json:"phase"`
		WinTeam string `json:"win_team"`
	} `json:"round"`
	Player struct {
		Steamid      string `json:"steamid"`
		Clan         string `json:"clan"`
		Name         string `json:"name"`
		ObserverSlot int    `json:"observer_slot"`
		Team         string `json:"team"`
		Activity     string `json:"activity"`
		State        struct {
			Health      int  `json:"health"`
			Armor       int  `json:"armor"`
			Helmet      bool `json:"helmet"`
			Defusekit   bool `json:"defusekit"`
			Flashed     int  `json:"flashed"`
			Smoked      int  `json:"smoked"`
			Burning     int  `json:"burning"`
			Money       int  `json:"money"`
			RoundKills  int  `json:"round_kills"`
			RoundKillhs int  `json:"round_killhs"`
			EquipValue  int  `json:"equip_value"`
		} `json:"state"`
	} `json:"player"`
	Allplayers struct {
		Num76561197997785085 struct {
			Clan         string `json:"clan"`
			Name         string `json:"name"`
			ObserverSlot int    `json:"observer_slot"`
			Team         string `json:"team"`
			State        struct {
				Health      int  `json:"health"`
				Armor       int  `json:"armor"`
				Helmet      bool `json:"helmet"`
				Flashed     int  `json:"flashed"`
				Burning     int  `json:"burning"`
				Money       int  `json:"money"`
				RoundKills  int  `json:"round_kills"`
				RoundKillhs int  `json:"round_killhs"`
				EquipValue  int  `json:"equip_value"`
			} `json:"state"`
			MatchStats struct {
				Kills   int `json:"kills"`
				Assists int `json:"assists"`
				Deaths  int `json:"deaths"`
				Mvps    int `json:"mvps"`
				Score   int `json:"score"`
			} `json:"match_stats"`
			Weapons struct {
			} `json:"weapons"`
			Position string `json:"position"`
		} `json:"76561197997785085"`
		Num76561198140309088 struct {
			Name         string `json:"name"`
			ObserverSlot int    `json:"observer_slot"`
			Team         string `json:"team"`
			State        struct {
				Health      int  `json:"health"`
				Armor       int  `json:"armor"`
				Helmet      bool `json:"helmet"`
				Flashed     int  `json:"flashed"`
				Burning     int  `json:"burning"`
				Money       int  `json:"money"`
				RoundKills  int  `json:"round_kills"`
				RoundKillhs int  `json:"round_killhs"`
				EquipValue  int  `json:"equip_value"`
			} `json:"state"`
			MatchStats struct {
				Kills   int `json:"kills"`
				Assists int `json:"assists"`
				Deaths  int `json:"deaths"`
				Mvps    int `json:"mvps"`
				Score   int `json:"score"`
			} `json:"match_stats"`
			Weapons struct {
			} `json:"weapons"`
			Position string `json:"position"`
		} `json:"76561198140309088"`
		Num76561198049042155 struct {
			Name         string `json:"name"`
			ObserverSlot int    `json:"observer_slot"`
			Team         string `json:"team"`
			State        struct {
				Health      int  `json:"health"`
				Armor       int  `json:"armor"`
				Helmet      bool `json:"helmet"`
				Defusekit   bool `json:"defusekit"`
				Flashed     int  `json:"flashed"`
				Burning     int  `json:"burning"`
				Money       int  `json:"money"`
				RoundKills  int  `json:"round_kills"`
				RoundKillhs int  `json:"round_killhs"`
				EquipValue  int  `json:"equip_value"`
			} `json:"state"`
			MatchStats struct {
				Kills   int `json:"kills"`
				Assists int `json:"assists"`
				Deaths  int `json:"deaths"`
				Mvps    int `json:"mvps"`
				Score   int `json:"score"`
			} `json:"match_stats"`
			Weapons struct {
				Weapon0 struct {
					Name     string `json:"name"`
					Paintkit string `json:"paintkit"`
					Type     string `json:"type"`
					State    string `json:"state"`
				} `json:"weapon_0"`
				Weapon1 struct {
					Name        string `json:"name"`
					Paintkit    string `json:"paintkit"`
					Type        string `json:"type"`
					AmmoReserve int    `json:"ammo_reserve"`
					State       string `json:"state"`
				} `json:"weapon_1"`
			} `json:"weapons"`
			Position string `json:"position"`
		} `json:"76561198049042155"`
		Num76561197960265755 struct {
			Name         string `json:"name"`
			ObserverSlot int    `json:"observer_slot"`
			Team         string `json:"team"`
			State        struct {
				Health      int  `json:"health"`
				Armor       int  `json:"armor"`
				Helmet      bool `json:"helmet"`
				Flashed     int  `json:"flashed"`
				Burning     int  `json:"burning"`
				Money       int  `json:"money"`
				RoundKills  int  `json:"round_kills"`
				RoundKillhs int  `json:"round_killhs"`
				EquipValue  int  `json:"equip_value"`
			} `json:"state"`
			MatchStats struct {
				Kills   int `json:"kills"`
				Assists int `json:"assists"`
				Deaths  int `json:"deaths"`
				Mvps    int `json:"mvps"`
				Score   int `json:"score"`
			} `json:"match_stats"`
			Weapons struct {
			} `json:"weapons"`
			Position string `json:"position"`
		} `json:"76561197960265755"`
		Num76561197979192443 struct {
			Clan         string `json:"clan"`
			Name         string `json:"name"`
			ObserverSlot int    `json:"observer_slot"`
			Team         string `json:"team"`
			State        struct {
				Health      int  `json:"health"`
				Armor       int  `json:"armor"`
				Helmet      bool `json:"helmet"`
				Defusekit   bool `json:"defusekit"`
				Flashed     int  `json:"flashed"`
				Burning     int  `json:"burning"`
				Money       int  `json:"money"`
				RoundKills  int  `json:"round_kills"`
				RoundKillhs int  `json:"round_killhs"`
				EquipValue  int  `json:"equip_value"`
			} `json:"state"`
			MatchStats struct {
				Kills   int `json:"kills"`
				Assists int `json:"assists"`
				Deaths  int `json:"deaths"`
				Mvps    int `json:"mvps"`
				Score   int `json:"score"`
			} `json:"match_stats"`
			Weapons struct {
				Weapon0 struct {
					Name     string `json:"name"`
					Paintkit string `json:"paintkit"`
					Type     string `json:"type"`
					State    string `json:"state"`
				} `json:"weapon_0"`
				Weapon1 struct {
					Name        string `json:"name"`
					Paintkit    string `json:"paintkit"`
					Type        string `json:"type"`
					AmmoClip    int    `json:"ammo_clip"`
					AmmoClipMax int    `json:"ammo_clip_max"`
					AmmoReserve int    `json:"ammo_reserve"`
					State       string `json:"state"`
				} `json:"weapon_1"`
				Weapon2 struct {
					Name        string `json:"name"`
					Paintkit    string `json:"paintkit"`
					Type        string `json:"type"`
					AmmoClip    int    `json:"ammo_clip"`
					AmmoClipMax int    `json:"ammo_clip_max"`
					AmmoReserve int    `json:"ammo_reserve"`
					State       string `json:"state"`
				} `json:"weapon_2"`
				Weapon3 struct {
					Name        string `json:"name"`
					Paintkit    string `json:"paintkit"`
					Type        string `json:"type"`
					AmmoReserve int    `json:"ammo_reserve"`
					State       string `json:"state"`
				} `json:"weapon_3"`
				Weapon4 struct {
					Name        string `json:"name"`
					Paintkit    string `json:"paintkit"`
					Type        string `json:"type"`
					AmmoReserve int    `json:"ammo_reserve"`
					State       string `json:"state"`
				} `json:"weapon_4"`
				Weapon5 struct {
					Name        string `json:"name"`
					Paintkit    string `json:"paintkit"`
					Type        string `json:"type"`
					AmmoReserve int    `json:"ammo_reserve"`
					State       string `json:"state"`
				} `json:"weapon_5"`
			} `json:"weapons"`
			Position string `json:"position"`
		} `json:"76561197979192443"`
		Num76561198158207405 struct {
			Name         string `json:"name"`
			ObserverSlot int    `json:"observer_slot"`
			Team         string `json:"team"`
			State        struct {
				Health      int  `json:"health"`
				Armor       int  `json:"armor"`
				Helmet      bool `json:"helmet"`
				Flashed     int  `json:"flashed"`
				Burning     int  `json:"burning"`
				Money       int  `json:"money"`
				RoundKills  int  `json:"round_kills"`
				RoundKillhs int  `json:"round_killhs"`
				EquipValue  int  `json:"equip_value"`
			} `json:"state"`
			MatchStats struct {
				Kills   int `json:"kills"`
				Assists int `json:"assists"`
				Deaths  int `json:"deaths"`
				Mvps    int `json:"mvps"`
				Score   int `json:"score"`
			} `json:"match_stats"`
			Weapons struct {
			} `json:"weapons"`
			Position string `json:"position"`
		} `json:"76561198158207405"`
		Num76561197960265756 struct {
			Name         string `json:"name"`
			ObserverSlot int    `json:"observer_slot"`
			Team         string `json:"team"`
			State        struct {
				Health      int  `json:"health"`
				Armor       int  `json:"armor"`
				Helmet      bool `json:"helmet"`
				Flashed     int  `json:"flashed"`
				Burning     int  `json:"burning"`
				Money       int  `json:"money"`
				RoundKills  int  `json:"round_kills"`
				RoundKillhs int  `json:"round_killhs"`
				EquipValue  int  `json:"equip_value"`
			} `json:"state"`
			MatchStats struct {
				Kills   int `json:"kills"`
				Assists int `json:"assists"`
				Deaths  int `json:"deaths"`
				Mvps    int `json:"mvps"`
				Score   int `json:"score"`
			} `json:"match_stats"`
			Weapons struct {
			} `json:"weapons"`
			Position string `json:"position"`
		} `json:"76561197960265756"`
		Num76561197960265754 struct {
			Name         string `json:"name"`
			ObserverSlot int    `json:"observer_slot"`
			Team         string `json:"team"`
			State        struct {
				Health      int  `json:"health"`
				Armor       int  `json:"armor"`
				Helmet      bool `json:"helmet"`
				Flashed     int  `json:"flashed"`
				Burning     int  `json:"burning"`
				Money       int  `json:"money"`
				RoundKills  int  `json:"round_kills"`
				RoundKillhs int  `json:"round_killhs"`
				EquipValue  int  `json:"equip_value"`
			} `json:"state"`
			MatchStats struct {
				Kills   int `json:"kills"`
				Assists int `json:"assists"`
				Deaths  int `json:"deaths"`
				Mvps    int `json:"mvps"`
				Score   int `json:"score"`
			} `json:"match_stats"`
			Weapons struct {
			} `json:"weapons"`
			Position string `json:"position"`
		} `json:"76561197960265754"`
	} `json:"allplayers"`
	PhaseCountdowns struct {
		Phase       string `json:"phase"`
		PhaseEndsIn string `json:"phase_ends_in"`
	} `json:"phase_countdowns"`
	Previously struct {
		PhaseCountdowns struct {
			PhaseEndsIn string `json:"phase_ends_in"`
		} `json:"phase_countdowns"`
	} `json:"previously"`
	Auth struct {
		Token string `json:"token"`
	} `json:"auth"`
}

func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		fmt.Println(string(body))
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
		var f Message
		err = json.Unmarshal(body, &f)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v", f)
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Printf("About to listen on 8080. Go to http://127.0.0.1:8080/")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	log.Fatal(err)
}
