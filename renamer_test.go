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
		out    string
		reason string
	}{
		{
			"Ghost_in_the_Shell_S.A.C_Individual_Eleven_[720p,BluRay,x264]_-_THORA",
			"Ghost.in.the.Shell.S.A.C.Individual.Eleven.720p.BluRay.x264.-.THORA",
			"replace underscore and comma, brackets",
		},
		{
			"There.Will.Be.Blood.2007.REPACK.720p.BluRay.x264-DON.chs&amp;eng[R3米粒修订].srt",
			"There.Will.Be.Blood.2007.REPACK.720p.BluRay.x264-DON.chs.amp.eng.R3米粒修订..srt",
			"replace &, commo",
		},
		{
			"Ghost_in_the_Shell_2_Innocence_(2004)_[720p,BluRay,DTS-ES,x264]_-_THORA",
			"Ghost.in.the.Shell.2.Innocence.2004.720p.BluRay.DTS-ES.x264.-.THORA",
			"replace parentheses",
		},
	}

	for _, input := range inputs {
		o := replaceOthers(input.src)
		if o != input.out {
			t.Errorf("%s:\ninput: %s\nexpected: %s\ngot: %s\n",
				input.reason, input.src, input.out, o,
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
