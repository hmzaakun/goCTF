package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net"
    "net/http"
    "sync"
	"strings"
	"strconv"
)
func extractLevel(s string) (string, error) {
	prefix := "Level:"
	if !strings.Contains(s, prefix) {
		return "", fmt.Errorf("level not found")
	}
	level := strings.TrimSpace(strings.TrimPrefix(s, prefix))
	return level, nil
}

func testPort(serverIP string, port int, wg *sync.WaitGroup) {
	var levelStr string
	var levelNum int
    defer wg.Done()
    address := fmt.Sprintf("%s:%d", serverIP, port)

    // Tentative de connexion au serveur
    conn, err := net.Dial("tcp", address)
    if err == nil {
        conn.Close()

        // Faire une requête HTTP GET pour /ping
        pingURL := fmt.Sprintf("http://%s:%d/ping", serverIP, port)
        respPing, err := http.Get(pingURL)
        if err == nil {
            defer respPing.Body.Close()
            bodyPing, _ := ioutil.ReadAll(respPing.Body)
            fmt.Printf("Port %d accessible - GET Response for /ping: %s\n", port, bodyPing)
        }

        // Effectuer une requête HTTP POST pour /signup
        signupURL := fmt.Sprintf("http://%s:%d/signup", serverIP, port)
        jsonStr := []byte(`{"User": "Hamza"}`)
        respSignup, err := http.Post(signupURL, "application/json", bytes.NewBuffer(jsonStr))
        if err == nil {
            defer respSignup.Body.Close()
            bodySignup, _ := ioutil.ReadAll(respSignup.Body)
            fmt.Printf("Port %d accessible - POST Response for /signup: %s\n", port, bodySignup)
        }

        // Effectuer une requête HTTP POST pour /check
        checkURL := fmt.Sprintf("http://%s:%d/check", serverIP, port)
        checkJsonStr := []byte(`{"User": "Hamza"}`)
        respCheck, err := http.Post(checkURL, "application/json", bytes.NewBuffer(checkJsonStr))
        if err == nil {
            defer respCheck.Body.Close()
            bodyCheck, _ := ioutil.ReadAll(respCheck.Body)
            fmt.Printf("Port %d accessible - POST Response for /check: %s\n", port, bodyCheck)
        }

		//26ac05af275df37b36a23a75dfb701c48ecbdb6a5a578674c36274737864f4ce
        // Effectuer une requête HTTP POST pour /secret
        secretURL := fmt.Sprintf("http://%s:%d/getUserSecret", serverIP, port)
        secretJsonStr := []byte(`{"User": "Hamza"}`)
        respSecret, err := http.Post(secretURL, "application/json", bytes.NewBuffer(secretJsonStr))
        if err == nil {
            defer respSecret.Body.Close()
            bodySecret, _ := ioutil.ReadAll(respSecret.Body)
            fmt.Printf("Port %d accessible - POST Response for /secret: %s\n", port, bodySecret)
        }

		// Effectuer une requête HTTP POST pour /getUserLevel
        getUserLevelURL := fmt.Sprintf("http://%s:%d/getUserLevel", serverIP, port)
        getUserLevelJsonStr := []byte(`{"User": "Hamza", "secret": "26ac05af275df37b36a23a75dfb701c48ecbdb6a5a578674c36274737864f4ce"}`)
        respGetUserLevel, err := http.Post(getUserLevelURL, "application/json", bytes.NewBuffer(getUserLevelJsonStr))
        if err == nil {
            defer respGetUserLevel.Body.Close()
            bodyGetUserLevel, _ := ioutil.ReadAll(respGetUserLevel.Body)
            fmt.Printf("Port %d accessible - POST Response for /getUserLevel: %s\n", port, bodyGetUserLevel)
			levelStr,err = extractLevel(string(bodyGetUserLevel))
			levelNum, err = strconv.Atoi(levelStr)
			levelNum = levelNum + 3
        }

		// Effectuer une requête HTTP POST pour /getUserPoints
        getUserPointsURL := fmt.Sprintf("http://%s:%d/getUserPoints", serverIP, port)
        getUserPointsJsonStr := []byte(`{"User": "Hamza", "secret": "26ac05af275df37b36a23a75dfb701c48ecbdb6a5a578674c36274737864f4ce"}`)
        respGetUserPoints, err := http.Post(getUserPointsURL, "application/json", bytes.NewBuffer(getUserPointsJsonStr))
        if err == nil {
            defer respGetUserPoints.Body.Close()
            bodyGetUserPoints, _ := ioutil.ReadAll(respGetUserPoints.Body)
            fmt.Printf("Port %d accessible - POST Response for /getUserPoints: %s\n", port, bodyGetUserPoints)
        }

		// Effectuer une requête HTTP POST pour /iNeedAHint
	    //morceau commenté car cet appel fait perdre des points au user mais ca marche
        // iNeedAHintURL := fmt.Sprintf("http://%s:%d/iNeedAHint", serverIP, port)
        // iNeedAHintJsonStr := []byte(`{"User": "Hamza", "secret": "26ac05af275df37b36a23a75dfb701c48ecbdb6a5a578674c36274737864f4ce"}`)
        // respINeedAHint, err := http.Post(iNeedAHintURL, "application/json", bytes.NewBuffer(iNeedAHintJsonStr))
        // if err == nil {
        //     defer respINeedAHint.Body.Close()
        //     bodyINeedAHint, _ := ioutil.ReadAll(respINeedAHint.Body)
        //     fmt.Printf("Port %d accessible - POST Response for /iNeedAHint: %s\n", port, bodyINeedAHint)
        // }

		// Effectuer une requête HTTP POST pour /enterChallenge
        getChallengeURL := fmt.Sprintf("http://%s:%d/enterChallenge", serverIP, port)
        getChallengeJsonStr := []byte(`{"User": "Hamza", "secret": "26ac05af275df37b36a23a75dfb701c48ecbdb6a5a578674c36274737864f4ce"}`)
        respGetChallenge, err := http.Post(getChallengeURL, "application/json", bytes.NewBuffer(getChallengeJsonStr))
        if err == nil {
            defer respGetChallenge.Body.Close()
            bodyGetChallenge, _ := ioutil.ReadAll(respGetChallenge.Body)
            fmt.Printf("Port %d accessible - POST Response for /getChallenge: %s\n", port, bodyGetChallenge)
        }

		// Effectuer une requête HTTP POST pour /submitSolution
        submitSolutionURL := fmt.Sprintf("http://%s:%d/submitSolution", serverIP, port)
        submitSolutionJsonStr := []byte(fmt.Sprintf(`{
            "User": "Hamza",
            "Secret": "26ac05af275df37b36a23a75dfb701c48ecbdb6a5a578674c36274737864f4ce",
            "Content": {
                "Level": %d,
                "Challenge": {
                    "Username": "Hamza",
                    "Secret": "26ac05af275df37b36a23a75dfb701c48ecbdb6a5a578674c36274737864f4ce",
                    "Points": 98
                },
                "Protocol": "MD5",
                "SecretKey": "Pasting code from the Internet into production code is like chewing gum found in the street."
            }
        }`,levelNum))
        respSubmitSolution, err := http.Post(submitSolutionURL, "application/json", bytes.NewBuffer(submitSolutionJsonStr))
        if err == nil {
            defer respSubmitSolution.Body.Close()
            bodySubmitSolution, _ := ioutil.ReadAll(respSubmitSolution.Body)
            fmt.Printf("Port %d accessible - POST Response for /submitSolution: %s\n", port, bodySubmitSolution)
        }
    }
}

func main() {
    serverIP := "10.49.122.144"
    minPort := 1000
    maxPort := 10000

    var wg sync.WaitGroup

    for port := minPort; port <= maxPort; port++ {
        wg.Add(1)
        go testPort(serverIP, port, &wg)
    }

    // Attendre que toutes les goroutines se terminent
    wg.Wait()
}
