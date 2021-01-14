package gopatchy

import "time"

type PatchDayObject struct {
	DocumentTitle struct {
		Value string `json:"Value"`
	} `json:"DocumentTitle"`
	DocumentType struct {
		Value string `json:"Value"`
	} `json:"DocumentType"`
	DocumentPublisher struct {
		ContactDetails struct {
			Value string `json:"Value"`
		} `json:"ContactDetails"`
		IssuingAuthority struct {
			Value string `json:"Value"`
		} `json:"IssuingAuthority"`
		Type int `json:"Type"`
	} `json:"DocumentPublisher"`
	DocumentTracking struct {
		Identification struct {
			ID struct {
				Value string `json:"Value"`
			} `json:"ID"`
			Alias struct {
				Value string `json:"Value"`
			} `json:"Alias"`
		} `json:"Identification"`
		Status          int    `json:"Status"`
		Version         string `json:"Version"`
		RevisionHistory []struct {
			Number      string    `json:"Number"`
			Date        time.Time `json:"Date"`
			Description struct {
				Value string `json:"Value"`
			} `json:"Description"`
		} `json:"RevisionHistory"`
		InitialReleaseDate time.Time `json:"InitialReleaseDate"`
		CurrentReleaseDate time.Time `json:"CurrentReleaseDate"`
	} `json:"DocumentTracking"`
	DocumentNotes []struct {
		Title    string `json:"Title"`
		Audience string `json:"Audience"`
		Type     int    `json:"Type"`
		Ordinal  string `json:"Ordinal"`
		Value    string `json:"Value"`
	} `json:"DocumentNotes"`
	ProductTree struct {
		Branch []struct {
			Items []struct {
				Items []struct {
					ProductID string `json:"ProductID"`
					Value     string `json:"Value"`
				} `json:"Items"`
				Type int    `json:"Type"`
				Name string `json:"Name"`
			} `json:"Items"`
			Type int    `json:"Type"`
			Name string `json:"Name"`
		} `json:"Branch"`
		FullProductName []struct {
			ProductID string `json:"ProductID"`
			Value     string `json:"Value"`
		} `json:"FullProductName"`
	} `json:"ProductTree"`
	Vulnerability []struct {
		Title struct {
			Value string `json:"Value"`
		} `json:"Title"`
		Notes []struct {
			Title   string `json:"Title"`
			Type    int    `json:"Type"`
			Ordinal string `json:"Ordinal"`
			Value   string `json:"Value,omitempty"`
		} `json:"Notes"`
		DiscoveryDateSpecified bool   `json:"DiscoveryDateSpecified"`
		ReleaseDateSpecified   bool   `json:"ReleaseDateSpecified"`
		CVE                    string `json:"CVE"`
		ProductStatuses        []struct {
			ProductID []string `json:"ProductID"`
			Type      int      `json:"Type"`
		} `json:"ProductStatuses"`
		Threats []struct {
			Description struct {
				Value string `json:"Value"`
			} `json:"Description"`
			ProductID     []string `json:"ProductID,omitempty"`
			Type          int      `json:"Type"`
			DateSpecified bool     `json:"DateSpecified"`
		} `json:"Threats"`
		CVSSScoreSets []struct {
			BaseScore     float64  `json:"BaseScore"`
			TemporalScore float64  `json:"TemporalScore"`
			Vector        string   `json:"Vector"`
			ProductID     []string `json:"ProductID"`
		} `json:"CVSSScoreSets"`
		Remediations    []interface{} `json:"Remediations"`
		Acknowledgments []struct {
			Name []struct {
				Value string `json:"Value"`
			} `json:"Name"`
			URL []string `json:"URL"`
		} `json:"Acknowledgments"`
		Ordinal         string `json:"Ordinal"`
		RevisionHistory []struct {
			Number      string    `json:"Number"`
			Date        time.Time `json:"Date"`
			Description struct {
				Value string `json:"Value"`
			} `json:"Description"`
		} `json:"RevisionHistory"`
	} `json:"Vulnerability"`
}
