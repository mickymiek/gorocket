package rest

import (
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
	"github.com/skilld-labs/gorocket/api"
)

type channelsResponse struct {
	Success  bool `json:"success"`
	Channels []api.Channel `json:"channels"`
}

type channelResponse struct {
	Success bool `json:"success"`
	Channel api.Channel `json:"channel"`
}

// Returns all channels that can be seen by the logged in user.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/list
func (c *Client) GetPublicChannels() ([]api.Channel, error) {
	request, _ := http.NewRequest("GET", c.getUrl() + "/api/v1/channels.list", nil)
	response := new(channelsResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Channels, nil
}

// Returns all channels that the user has joined.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/list-joined
func (c *Client) GetJoinedChannels() ([]api.Channel, error) {
	request, _ := http.NewRequest("GET", c.getUrl() + "/api/v1/channels.list.joined", nil)
	response := new(channelsResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Channels, nil
}

// Returns all groups that the user has joined.
//
// https://rocket.chat/docs/developer-guides/rest-api/groups/list
func (c *Client) GetJoinedGroups() ([]api.Channel, error) {
	request, _ := http.NewRequest("GET", c.getUrl() + "/api/v1/groups.list", nil)
	response := new(channelsResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Groups, nil
}

// Joins a channel. The id of the channel has to be not nil.
//
// This function is not supported by the current Client.Chat release version 0.48.2.
func (c *Client) JoinChannel(channel *api.Channel) error {
	var body = fmt.Sprintf(`{ "roomId": "%s" }`, channel.Id)
	request, _ := http.NewRequest("POST", c.getUrl() + "/api/v1/channels.join", bytes.NewBufferString(body))
	return c.doRequest(request, new(statusResponse))
}

// Creates a group with users. The username(s) needs to be registered in RC.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/create
func (c *Client) CreateGroup(channel *api.Channel) error {
	u, err := json.Marshal(channel.UserNames)
	if err != nil {
		return err
	}
	var body = fmt.Sprintf(`{ "name": "%s", "members": %s }`, channel.Name, u)
	request, _ := http.NewRequest("POST", c.getUrl() + "/api/v1/groups.create", bytes.NewBufferString(body))
	return c.doRequest(request, new(statusResponse))
}

// Archives a channel. The roomId has to be set.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/archive
func (c *Client) ArchiveGroup(channel *api.Channel) error {
	var body = fmt.Sprintf(`{ "roomId": "%s" }`, channel.Id)
	request, _ := http.NewRequest("POST", c.getUrl() + "/api/v1/groups.archive", bytes.NewBufferString(body))
	return c.doRequest(request, new(statusResponse))
}

// Leaves a channel. The id of the channel has to be not nil.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/leave
func (c *Client) LeaveChannel(channel *api.Channel) error {
	var body = fmt.Sprintf(`{ "roomId": "%s"}`, channel.Id)
	request, _ := http.NewRequest("POST", c.getUrl() + "/api/v1/channels.leave", bytes.NewBufferString(body))
	return c.doRequest(request, new(statusResponse))
}

// Get information about a channel. That might be useful to update the usernames.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/info
func (c *Client) GetChannelInfo(channel *api.Channel) (*api.Channel, error) {
	var url = fmt.Sprintf("%s/api/v1/channels.info?roomId=%s", c.getUrl(), channel.Id)
	request, _ := http.NewRequest("GET", url, nil)
	response := new(channelResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return &response.Channel, nil
}
