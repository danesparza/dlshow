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
	filename := "Once.Upon.a.Time.S03E01.720p.HDTV.X264-DIMENSION.mkv"

	//	Act
	showInfo, err := dlshow.GetEpisodeInfo(filename)

	//	Assert
	if err != nil {
		t.Errorf("Should have returned show info: %s", filename)
	}

	if showInfo.ParseType != dlshow.ParseTypeSE {
		t.Errorf("Incorrect parse type: %v", showInfo.ParseType)
	}
	if showInfo.ShowName != "Once Upon a Time" {
		t.Errorf("Incorrect show name parsed: %v", showInfo.ShowName)
	}
	if showInfo.SeasonNumber != 3 {
		t.Errorf("Expected season 3 but got %v instead", showInfo.SeasonNumber)
	}

	if showInfo.EpisodeNumber != 1 {
		t.Errorf("Expected episode 1 but got %v instead", showInfo.EpisodeNumber)
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
