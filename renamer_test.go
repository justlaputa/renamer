package main

import "testing"

func TestReplaceSpace(t *testing.T) {
	inputs := []struct {
		src    string
		reason string
	}{
		{"Straw Dogs 1971 720p BluRay FLAC1.0 x264-DON", "replace space to dot"},
		{"Straw  Dogs 1971 720p BluRay FLAC1.0 x264-DON", "replace continous space to single dot"},
		{" Straw Dogs 1971 720p BluRay FLAC1.0 x264-DON", "remove prefix space"},
		{"Straw Dogs 1971 720p BluRay FLAC1.0 x264-DON ", "remove trailing space"},
	}

	output := "Straw.Dogs.1971.720p.BluRay.FLAC1.0.x264-DON"

	for _, input := range inputs {
		o := replaceSpece(input.src)
		if o != output {
			t.Errorf("%s:\ninput: %s\nexpected: %s\ngot: %s\n",
				input.reason, input.src, output, o,
			)
		}
	}
}

func TestReplaceOthers(t *testing.T) {
	inputs := []struct {
		src    string
		reason string
	}{
		{"Ghost_in_the_Shell_S.A.C_Individual_Eleven_[720p,BluRay,x264]_-_THORA", "replace underscore and comma"},
	}

	output := "Ghost.in.the.Shell.S.A.C.Individual.Eleven.720p.BluRay.x264.-.THORA"
	for _, input := range inputs {
		o := replaceOthers(input.src)
		if o != output {
			t.Errorf("%s:\ninput: %s\nexpected: %s\ngot: %s\n",
				input.reason, input.src, output, o,
			)
		}
	}
}

func TestDeduplicate(t *testing.T) {
	inputs := []struct {
		src    string
		reason string
	}{
		{"Ghost.in.the.Shell.S.A.C.Individual.Eleven.720p.BluRay.x264.-.THORA", "replace dot connect hyphen"},
		{"Ghost..in.the.Shell.S.A.C.Individual.Eleven.720p.BluRay.x264-THORA", "remove double dot"},
		{"Ghost.in.the.Shell.S.A.C.Individual.Eleven.720p.BluRay...x264-THORA", "remove tripple dot"},
		{"Ghost.in.the.Shell.S.A.C.Individual.Eleven.720p.BluRay.x264.-THORA", "replace dot hyphen"},
		{"Ghost.in.the.Shell.S.A.C.Individual.Eleven.720p.BluRay.x264-.THORA", "replace hyphen dot"},
	}

	output := "Ghost.in.the.Shell.S.A.C.Individual.Eleven.720p.BluRay.x264-THORA"
	for _, input := range inputs {
		o := deduplicate(input.src)
		if o != output {
			t.Errorf("%s:\ninput:  \t%s\nexpected:\t%s\nacctual:\t%s\n",
				input.reason, input.src, output, o,
			)
		}
	}
}