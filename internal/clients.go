package internal

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type client_session struct {
	Host   string
	Secret string
	Token  string

	access_mutex sync.Mutex
	LastAccess   time.Time
	Connections  []*vcenter_connection
}

// Slice containing our active clients
var clients []client_session
var clients_mutex sync.Mutex

func checkClientConnections() {
	clients_mutex.Lock()

	// From:
	// https://stackoverflow.com/questions/20545743/how-to-remove-items-from-a-slice-while-ranging-over-it
	// This leaks memory until the slice is garbage collected ... so not ideal.

	index := 0
	for i := 0; i < len(clients); i++ {
		c := &clients[i]

		// check if the c.LastAccess + timeout value is later than now
		interval := time.Since(c.LastAccess)
		if interval.Minutes() > 60 {
			// if so, remove from the clients slice
			log.Printf("Client connection from host %v with token %v has timed-out due to inactivity.\n", c.Host, c.Token)
			continue

		}

		clients[index] = *c
		index++
	}
	clients = clients[:index]

	clients_mutex.Unlock()
}

func parseRequestAddress(host string) string {
	// Parse the source IP
	var result string
	index := strings.LastIndex(host, ":")
	if index > -1 {
		result = host[0:index]
	}
	return result
}

func registerNewSession(addr string, secret string) (client *client_session, err error) {
	host := parseRequestAddress(addr)

	clients_mutex.Lock()

	for i := 0; i < len(clients); i++ {
		if clients[i].Host == host && clients[i].Secret == secret {
			c := &clients[i]
			clients_mutex.Unlock()
			return c, nil
		}
	}

	// Generate the token. This could fail.
	id, err := uuid.NewRandom()
	if err != nil {
		clients_mutex.Unlock()
		return nil, err
	}

	// register new session and return the token
	var cl client_session
	cl.Host = host
	cl.Secret = secret
	cl.Token = id.String()
	cl.LastAccess = time.Now()

	fmt.Println("Host " + host + " connected.")

	clients = append(clients, cl)
	clients_mutex.Unlock()
	return &cl, nil
}

func verifyClientAcess(addr string, token string) (client *client_session, err error) {
	host := parseRequestAddress(addr)

	clients_mutex.Lock()
	for i := 0; i < len(clients); i++ {
		if clients[i].Host == host && clients[i].Token == token {
			clients[i].LastAccess = time.Now()
			c := &clients[i]
			clients_mutex.Unlock()
			return c, nil
		}
	}
	clients_mutex.Unlock()

	return nil, errors.New("no client access")
}

func registerClientvCenterConnection(client *client_session, vc_url, username string) *vcenter_connection {
	var result *vcenter_connection

	client.access_mutex.Lock()

	found := false
	for i := 0; i < len(client.Connections); i++ {
		vc := client.Connections[i]
		if vc_url == vc.URL && username == vc.Username {
			found = true
			result = vc
			break
		}
	}

	if !found {
		connections_mutex.Lock()

		global_found := false
		for i := 0; i < len(connections); i++ {
			var vc *vcenter_connection
			vc = connections[i]
			if vc.URL == vc_url && vc.Username == username {
				global_found = true
				result = vc
				client.Connections = append(client.Connections, vc)
				break
			}
		}

		if !global_found {
			var vc *vcenter_connection
			vc = new(vcenter_connection)
			vc.URL = vc_url
			vc.Username = username

			parsed_url, err := url.Parse(vc_url)
			if err == nil {
				base_name := parsed_url.Hostname()
				desired_name := base_name
				base_name += "_"

				/*
					We make sure the name of the connection is unique
				*/
				number := 1
				for {
					// See if we can find desired name
					found := false
					for i := 0; i < len(connections); i++ {
						conn := connections[i]
						if conn.Name == desired_name {
							found = true
							break
						}
					}

					if !found {
						break
					}

					// Change desired name to something unique and go again
					desired_name = base_name + fmt.Sprint(number)
					number++
				}

				vc.Name = desired_name
			}

			connections = append(connections, vc)
			client.Connections = append(client.Connections, vc)
			result = vc
		}

		connections_mutex.Unlock()
	}

	client.access_mutex.Unlock()
	return result
}

func getClientvCenterConnectionByName(client *client_session, vc_name string) *vcenter_connection {
	var result *vcenter_connection

	client.access_mutex.Lock()

	for i := 0; i < cap(client.Connections); i++ {
		vc := client.Connections[i]
		if vc_name == vc.Name {
			result = vc
			break
		}
	}

	client.access_mutex.Unlock()

	return result
}
