package gopngchunkschara

// CharacterCardV1V2 - https://github.com/malfoyslastname/character-card-spec-v2/blob/main/spec_v2.md
type CharacterCardV1V2 struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Personality    string `json:"personality"`
	Scenario       string `json:"scenario"`
	FirstMes       string `json:"first_mes"`
	MesExample     string `json:"mes_example"`
	Creatorcomment string `json:"creatorcomment"`
	Avatar         string `json:"avatar"`
	Chat           string `json:"chat"`
	Talkativeness  string `json:"talkativeness"`
	Fav            bool   `json:"fav"`
	CreateDate     string `json:"create_date"`
	Spec           string `json:"spec"`
	SpecVersion    string `json:"spec_version"`
	Data           struct {
		Name                    string        `json:"name"`
		Description             string        `json:"description"`
		Personality             string        `json:"personality"`
		Scenario                string        `json:"scenario"`
		FirstMes                string        `json:"first_mes"`
		MesExample              string        `json:"mes_example"`
		CreatorNotes            string        `json:"creator_notes"`
		SystemPrompt            string        `json:"system_prompt"`
		PostHistoryInstructions string        `json:"post_history_instructions"`
		Tags                    []interface{} `json:"tags"`
		Creator                 string        `json:"creator"`
		CharacterVersion        string        `json:"character_version"`
		AlternateGreetings      []interface{} `json:"alternate_greetings"`
		Extensions              struct {
			Talkativeness string `json:"talkativeness"`
			Fav           bool   `json:"fav"`
		} `json:"extensions"`
	} `json:"data"`
}
