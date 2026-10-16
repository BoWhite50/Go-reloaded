package main

import (
	"go-reloaded/tools"
	"testing"
)

func TestAssemblage(t *testing.T) {
	tests := []struct {
		entree  string
		attendu string
	}{
		{"it (cap) was the best of times, it was the worst of times (up) , it was the age of foolishness (cap, 6) , IT WAS THE (low, 3) winter of despair.",
			"It was the best of times, it was the worst of TIMES, It Was The Age Of Foolishness, it was the winter of despair."},
		{"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},
		{"There is no greater agony than bearing a untold story inside you.", "There is no greater agony than bearing an untold story inside you."},
		{"Punctuation tests are ... kinda boring ,what do you think ?", "Punctuation tests are... kinda boring, what do you think?"},
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"Ready, set, go (up) !", "Ready, set, GO!"},
		{"", ""},
		// bugs corrigés : on garde un test pour être sûr qu'ils ne reviennent pas
		{"(up) (low) hello", "hello"},
		{"(up, 2) (cap) hello world", "hello world"},
		{"' hello . '", "'hello.'"},
		{"a ' amazing ' day", "an 'amazing' day"},
		{`He said: " hello world " and ' bye '`, `He said: "hello world" and 'bye'`},
	}

	for _, test := range tests {
		resultat := tools.Assemblage([]string{test.entree})[0]
		if resultat != test.attendu {
			t.Errorf("entrée : %q\n attendu : %q\n obtenu : %q", test.entree, test.attendu, resultat)
		}
	}
}

// TestTextesInventes donne à Assemblage des phrases inventées
// qui mélangent plusieurs règles en même temps.
func TestTextesInventes(t *testing.T) {
	tests := []struct {
		entree  string
		attendu string
	}{
		{"my cat ate 3 (hex) fish today , and she wants more (up) !",
			"my cat ate 3 fish today, and she wants MORE!"},
		{"I found a old map in the attic . it led to a island (cap) !?",
			"I found an old map in the attic. it led to an Island!?"},
		{"the robot said : ' hello human ' and then it stopped working ...",
			"the robot said: 'hello human' and then it stopped working..."},
		{"we have 1010 (bin) days left before the exam (up, 2) , don't panic (cap) .",
			"we have 10 days left before THE EXAM, don't Panic."},
		{"my teacher is VERY STRICT (low, 2) but a honest person ; she always says ' work hard ' .",
			"my teacher is very strict but an honest person; she always says 'work hard'."},
		{"the pizza costs FF (hex) cents , that is a lot of money (cap, 3) ?",
			"the pizza costs 255 cents, that is a Lot Of Money?"},
	}

	for _, test := range tests {
		resultat := tools.Assemblage([]string{test.entree})[0]
		if resultat != test.attendu {
			t.Errorf("entrée : %q\n attendu : %q\n obtenu : %q", test.entree, test.attendu, resultat)
		}
	}
}
