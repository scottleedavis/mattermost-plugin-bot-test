package main

import (
	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
	"net/http"
)

func (p *Plugin) InitAPI() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/button", p.handleButton).Methods("POST")
	r.HandleFunc("/button2", p.handleButton2).Methods("POST")
	r.HandleFunc("/button-ephemeral", p.handleButtonEphemeral).Methods("POST")

	return r
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
}

func (p *Plugin) handleButton(w http.ResponseWriter, r *http.Request) {

	request := model.PostActionIntegrationRequestFromJson(r.Body)

	p.API.LogInfo("handleButton")
	if post, pErr := p.API.GetPost(request.PostId); pErr != nil {
		p.API.LogError("unable to get post " + pErr.Error())
		writePostActionIntegrationResponseError(w, &model.PostActionIntegrationResponse{})
	} else {
		post.Props = model.StringInterface{}
		post.Message = "clicked button"
		p.API.UpdatePost(post)

	}
	writePostActionIntegrationResponseOk(w, &model.PostActionIntegrationResponse{})

}

func (p *Plugin) handleButton2(w http.ResponseWriter, r *http.Request) {

	request := model.PostActionIntegrationRequestFromJson(r.Body)

	p.API.LogInfo("handleButton2")
	if post, pErr := p.API.GetPost(request.PostId); pErr != nil {
		p.API.LogError("unable to get post " + pErr.Error())
		writePostActionIntegrationResponseError(w, &model.PostActionIntegrationResponse{})
	} else {
		post.Props = model.StringInterface{}
		post.Message = "clicked button2"
		p.API.UpdatePost(post)
	}
	writePostActionIntegrationResponseOk(w, &model.PostActionIntegrationResponse{})

}

func (p *Plugin) handleButtonEphemeral(w http.ResponseWriter, r *http.Request) {

	request := model.PostActionIntegrationRequestFromJson(r.Body)

	post := &model.Post{
		Id:        request.PostId,
		UserId:    request.UserId,
		ChannelId: request.ChannelId,
		Message:   "clicked button ephemeral",
	}
	p.API.UpdateEphemeralPost(request.UserId, post)

	writePostActionIntegrationResponseOk(w, &model.PostActionIntegrationResponse{})

}

func writePostActionIntegrationResponseOk(w http.ResponseWriter, response *model.PostActionIntegrationResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response.ToJson())
}

func writePostActionIntegrationResponseError(w http.ResponseWriter, response *model.PostActionIntegrationResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write(response.ToJson())
}
