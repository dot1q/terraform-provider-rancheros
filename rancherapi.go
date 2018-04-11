package main

import (
	"fmt"
	"net/http"
	"crypto/tls"
	"strconv"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
)

// Ignore invalid certs
var transport = &http.Transport {
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}

// Set http client
var client = &http.Client {
	Transport: transport,
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func genApiKey(obj *ApiKeyDescriptor) (_err error) {
	var jsonStr = []byte(fmt.Sprintf(`{"name":"%s","description":"%s"}`, obj.Name, obj.Description))

	// go ahead and create the api key
	req, _err := http.NewRequest("POST", fmt.Sprintf("%s/v2-beta/projects/%s/apiKeys", obj.Host, obj.ProjectId), bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
        if obj.AccessKey != "" && obj.SecretKey != "" {
		req.Header.Add("Authorization","Basic " + basicAuth(obj.AccessKey,obj.SecretKey))
        }

	resp, _err := client.Do(req)
	if _err != nil {
		return _err
	}

	if resp.StatusCode != 201 {
		return fmt.Errorf("Unhandled HTTP code %s for genereating an API key", strconv.Itoa(resp.StatusCode))
	}

	// close the body stream
	defer resp.Body.Close()

	// make sure the json can parse
	out, _err := ioutil.ReadAll(resp.Body)

	if _err != nil {
		return _err
	}

	var jsonResult map[string]interface{}
	json.Unmarshal([]byte(out), &jsonResult)
	//json := jsonResult["publicValue"].(string)
	obj.UUID	 = jsonResult["id"].(string)
	obj.GenAccessKey = jsonResult["publicValue"].(string)
	obj.GenSecretKey = jsonResult["secretValue"].(string)

	return
}

func readApiKey(obj *ApiKeyDescriptor) (_err error) {
        var jsonStr = []byte(fmt.Sprintf(`{"name":"%s","description":"%s"}`, obj.Name, obj.Description))

        // go ahead and create the api key
	req, _err := http.NewRequest("GET", fmt.Sprintf("%s/v2-beta/projects/%s/apiKeys/%s", obj.Host, obj.ProjectId, obj.UUID), bytes.NewBuffer(jsonStr))
        req.Header.Set("Content-Type", "application/json")
        if obj.GenAccessKey != "" && obj.GenSecretKey != "" {
                req.Header.Add("Authorization","Basic " + basicAuth(obj.GenAccessKey,obj.GenSecretKey))
        }

	resp, _err := client.Do(req)
        if _err != nil {
                return _err
        }

        if resp.StatusCode != 200 {
                return fmt.Errorf("Unhandled HTTP code %s for reading an API key", strconv.Itoa(resp.StatusCode))
        }

        // close the body stream
        defer resp.Body.Close()

        // make sure the json can parse
        out, _err := ioutil.ReadAll(resp.Body)

        if _err != nil {
                return _err
        }

        var jsonResult map[string]interface{}
        json.Unmarshal([]byte(out), &jsonResult)
        //json := jsonResult["publicValue"].(string)
        obj.UUID         = jsonResult["id"].(string)
	obj.Name	 = jsonResult["name"].(string)
	obj.Description  = jsonResult["description"].(string)
//        obj.GenAccessKey = jsonResult["publicValue"].(string)
//        obj.GenSecretKey = jsonResult["secretValue"].(string)

        return
}

func updateApiKey(obj *ApiKeyDescriptor) (_err error) {
	var jsonStr = []byte(fmt.Sprintf(`{"name":"%s","description":"%s"}`, obj.Name, obj.Description))

	// go ahead and update the api key
	req, _err := http.NewRequest("PUT", fmt.Sprintf("%s/v2-beta/projects/%s/apiKeys/%s", obj.Host,obj.ProjectId, obj.UUID), bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if obj.GenAccessKey != "" && obj.GenSecretKey != "" {
		req.Header.Add("Authorization","Basic " + basicAuth(obj.GenAccessKey,obj.GenSecretKey))
	}

	resp, _err := client.Do(req)
	if _err != nil {
		return _err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("Unhandled HTTP code %s for deleting an API key", strconv.Itoa(resp.StatusCode))
	}

	return nil
}

func delApiKey(obj *ApiKeyDescriptor) (_err error) {
        // go ahead and create the api key
	req, _err := http.NewRequest("POST", fmt.Sprintf("%s/v2-beta/projects/%s/apiKeys/%s?action=deactivate", obj.Host, obj.ProjectId, obj.UUID), nil)
        req.Header.Set("Content-Type", "application/json")
        if obj.GenAccessKey != "" && obj.GenSecretKey != "" {
                req.Header.Add("Authorization","Basic " + basicAuth(obj.GenAccessKey,obj.GenSecretKey))
        }

	resp, _err := client.Do(req)
        if _err != nil {
                return _err
        }

        if resp.StatusCode != 404 && resp.StatusCode != 200 && resp.StatusCode != 202 {
                return fmt.Errorf("Unhandled HTTP code %s for deleting an API key", strconv.Itoa(resp.StatusCode))
	}

	return nil
}
