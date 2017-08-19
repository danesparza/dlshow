package dlshow_test

import (
	"testing"

	"github.com/danesparza/dlshow"
)

//	GetEpisodeInfo should return an error if it
//	isn't passed a valid filename
func TestGetEpisodeInfo_InvalidFilename_ReturnsError(t *testing.T) {

	//	Arrange
	filename := "c:\\temp\\"

	//	Act
	_, err := dlshow.GetEpisodeInfo(filename)

	//	Assert
	if err == nil {
		t.Errorf("Should have returned an error for the filename: %s", filename)
	}
}

func TestGetEpisodeInfo_ValidFilename_ReturnsEpisodeInfo(t *testing.T) {

	//	Arrange
	episodeTests := []struct {
		filename              string
		expectedShowName      string
		expectedSeasonNumber  int
		expectedEpisodeNumber int
	}{
		//	Episode tests
		{"Once.Upon.a.Time.S03E01.720p.HDTV.X264-DIMENSION.mkv", "Once Upon a Time", 3, 1},
		{"The.Big.Bang.Theory.S01E17.720p.HDTV.X264-MRSK.mkv", "The Big Bang Theory", 1, 17},
	}

	for _, test := range episodeTests {

		//	Act
		showInfo, err := dlshow.GetEpisodeInfo(test.filename)

		//	Assert
		if err != nil {
			t.Errorf("Should have returned show info: %s", test.filename)
		}

		if showInfo.ShowName != test.expectedShowName {
			t.Errorf("Expected show %v but got %v instead", test.expectedShowName, showInfo.ShowName)
		}

		if showInfo.SeasonNumber != test.expectedSeasonNumber {
			t.Errorf("Expected season %v but got %v instead", test.expectedSeasonNumber, showInfo.SeasonNumber)
		}

		if showInfo.EpisodeNumber != test.expectedEpisodeNumber {
			t.Errorf("Expected episode %v but got %v instead", test.expectedEpisodeNumber, showInfo.EpisodeNumber)
		}
	}
}

func TestGetEpisodeInfo_ValidAirdateFilename_ReturnsEpisodeInfo(t *testing.T) {
	//	Arrange
	filename := "Colbert.Report.2013.10.10.Reed.Albergotti.and.Vanessa.OConnell.HDTV.x264-LMAO.mp4"

	//	Act
	showInfo, err := dlshow.GetEpisodeInfo(filename)

	//	Assert
	if err != nil {
		t.Errorf("Should have returned show info: %s", filename)
	}

	if showInfo.ParseType != dlshow.ParseTypeDate {
		t.Errorf("Incorrect parse type: %v", showInfo.ParseType)
	}
	if showInfo.ShowName != "Colbert Report" {
		t.Errorf("Incorrect show name parsed: %v", showInfo.ShowName)
	}
	if showInfo.AiredYear != 2013 {
		t.Errorf("Expected year 2013 but got %v instead", showInfo.AiredYear)
	}

	if showInfo.AiredMonth != 10 {
		t.Errorf("Expected month 10 but got %v instead", showInfo.AiredMonth)
	}
}
