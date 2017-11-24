package main

import (
	"fmt"
	"net/http"
	"math/rand"
	"regexp"
	"time"
)

func elizaResponse(w http.ResponseWriter, r *http.Request) {

	input := r.URL.Query().Get("value")

	// All Elizas Responses
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		responses := []string {
			"Why don't you tell me more about your father?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\bmother\b.*`, input); matched {
		responses := []string {
			"Why don't you tell me more about your mother?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\bbrother\b.*`, input); matched {
		responses := []string {
			"Why don't you tell me more about your brother?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\bsister\b.*`, input); matched {
		responses := []string {
			"Why don't you tell me more about your sister?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Hh]ello|[Hh]ey|[Hh]i\b.*`, input); matched{
		responses := []string {
			"Hello, how are you?",
			"Howaya",
			"What's up?",
			"Oh hey",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	} else 

	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]hy\b|why.*`, input); matched{
		responses := []string {
			"Why do you ask that?",
			"I can't remember",
			"Who knows?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]ho.*`, input); matched{
		responses := []string {
			"I'm not sure, any other news?",
			"Can't remember them",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]hat.*`, input); matched{
		responses := []string {
			"I'm not sure actually",
			"You know",
			".........",
			"What? Who? How?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]hen.*`, input); matched{
		responses := []string {
			"I'm not sure",
			"How long ago was that?",
			
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]here.*`, input); matched{
		responses := []string {
			"You'll have to look on Google Maps",
			"I couldn't tell you",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Hh]ow.*`, input); matched{
		responses := []string {
			"Thats how",
			"Don't know, don't care",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*I am|I'm*`, input); matched{
		responses := []string {
			"How so?",
			"Why are you like that?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}
	
	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]as I*`, input); matched{
		responses := []string {
			"What were you thinking?",
			"Were you?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\bI was\b*`, input); matched{
		responses := []string {
			"Were you really?",
			"Why do you tell me that you were?",
			"Were you?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\bI am|I'm\b*`, input); matched{
		responses := []string {
			"Are you really?",
			"Pass no heed on me",
			"Why is that?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\bI do\b*`, input); matched{
		responses := []string {
			"I'm glad",
			"You do what? My memory isn't great",
			"Why do you?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}
	if matched, _ := regexp.MatchString(`(?i).*\bI dont|I don't\b*`, input); matched{
		responses := []string {
			"I'm not surprised",
			"Why don't you?",
			"Howcome?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Pp]erhaps\b*`, input); matched{
		responses := []string {
			"Indeed",
			"Why aren't you more sure?",
			"Nothing a google search won't solve",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}
	if matched, _ := regexp.MatchString(`(?i).*you are|you're*`, input); matched{
		responses := []string {
			"You don't know me",
			"I'm Eliza, actually",
			"Yes I am",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Yy]ou\b*`, input); matched{
		responses := []string {
			"I don't like the attention being on me",
			"Ask me again when you get to know me",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Yy]es\b*`, input); matched{
		responses := []string {
			"Are you certain?",
			"You're positive about it",
			"I see",
			"Indeed",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Nn]o\b*`, input); matched{
		responses := []string {
			"Are you saying no just to be negative?",
			"Why not yes?",
			"I hope you're happy",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Bb]ecause\b*`, input); matched{
		responses := []string {
			"Is that the only reason?",
			"Tell me the truth",
			"That's because?",
			"What other reasons might there be?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]hy don't you\b*`, input); matched{
		responses := []string {
			"That's my business",
			"I'm a computer program",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ll]ike\b*`, input); matched{
		responses := []string {
			"I would like that too",
			"I'm glad",
			"Oh wow",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ll]ove\b*`, input); matched{
		responses := []string {
			"I'm flattered",
			"T'm so happy for you",
			"Unfortunately I am not real",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\bI would\b*`, input); matched{
		responses := []string {
			"Really? Why would you?",
			"Why would you?",
			"You're telling me you would? Interesting. Go on.",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Hh]ate\b*`, input); matched{
		responses := []string {
			"Howcome?",
			"Oh no",
			"Since when?",
			"I don't feel emotion and even I think that's hard",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ii]t is\b*`, input); matched{
		responses := []string {
			"Is it really?",
			"Who knew!",
			"Oh is it?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return

	} else {
		responses := []string {

			"Lets talk about something else",
			"Tell me more about you",
			"I'm not really sure what you mean",
			"Can you re-word that please",
			"Is there anything else you wanna talk about?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}
	
}

func main() {

	// Seed Time for Random Response
	rand.Seed(time.Now().Unix())

	// Gets File From Static Folder
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Handles the Function elizaResponse
	http.HandleFunc("/elizaResponse", elizaResponse)

	// Enter This After localhost
	http.ListenAndServe(":8080", nil)
}