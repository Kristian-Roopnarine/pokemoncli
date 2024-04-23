package pokeapi

const PokemonApiUrl = "pokemon"

type PokemonResponse struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool             `json:"is_hidden"`
		Slot     int              `json:"slot"`
		Ability  NamedAPIResource `json:"ability"`
	} `json:"abilities"`
	Forms       []NamedAPIResource `json:"forms"`
	GameIndices []struct {
		GameIndex int              `json:"game_index"`
		Version   NamedAPIResource `json:"version"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item           NamedAPIResource `json:"item"`
		VersionDetails struct {
			Rarity  int              `json:"rarity"`
			Version NamedAPIResource `json:"version"`
		}
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move                NamedAPIResource `json:"move"`
		VersionGroupDetails struct {
			LevelLearnedAt  int              `json:"level_learned_at"`
			VersionGroup    NamedAPIResource `json:"version_group"`
			MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
		}
	} `json:"moves"`
	Species NamedAPIResource `json:"species"`
	Sprites struct {
		BackDefault      string  `json:"back_default"`
		BackFemale       *string `json:"back_female"`
		BackShiny        string  `json:"back_shiny"`
		BackShinyFemale  *string `json:"back_shiny_female"`
		FrontDefault     string  `json:"front_default"`
		FrontFemale      *string `json:"front_female"`
		FrontShiny       string  `json:"front_shiny"`
		FrontShinyFemale *string `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string  `json:"front_default"`
				FrontFemale  *string `json:"front_female"`
			}
			Home struct {
				FrontDefault     string  `json:"front_default"`
				FrontFemale      *string `json:"front_female"`
				FrontShiny       string  `json:"front_shiny"`
				FrontShinyFemale *string `json:"front_shiny_female"`
			}
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackDefault      string  `json:"back_default"`
				BackFemale       *string `json:"back_female"`
				BackShiny        string  `json:"back_shiny"`
				BackShinyFemale  *string `json:"back_shiny_female"`
				FrontDefault     string  `json:"front_default"`
				FrontFemale      *string `json:"front_female"`
				FrontShiny       string  `json:"front_shiny"`
				FrontShinyFemale *string `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
		} `json:"versions"`
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		BaseStat int              `json:"base_stat"`
		Effort   int              `json:"effort"`
		Stat     NamedAPIResource `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int              `json:"slot"`
		Type NamedAPIResource `json:"type"`
	} `json:"types"`
	PastTypes []struct {
		Generation NamedAPIResource `json:"generation"`
		Types      []struct {
			Slot int              `json:"slot"`
			Type NamedAPIResource `json:"type"`
		}
	} `json:"past_types"`
}
