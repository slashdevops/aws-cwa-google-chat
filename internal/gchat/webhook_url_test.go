package gchat

import (
	"net/url"
	"reflect"
	"testing"
)

func TestNewWebhookURL(t *testing.T) {
	type args struct {
		u *url.URL
	}
	tests := []struct {
		name    string
		args    args
		want    *WebhookURL
		wantErr bool
	}{
		{
			name: "valid URL",
			args: args{
				u: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/v1/spaces/spaceID/messages/",
					RawQuery: "key=key&token=token",
				},
			},
			want: &WebhookURL{
				URL: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/v1/spaces/spaceID/messages/",
					RawQuery: "key=key&token=token",
				},
				SpaceID:    "spaceID",
				APIVersion: "v1",
				Key:        "key",
				Token:      "token",
			},
			wantErr: false,
		},
		{
			name: "invalid URL, invalid path",
			args: args{
				u: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/wherever",
					RawQuery: "key=key&token=token",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid URL, unsupported API Version",
			args: args{
				u: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/badpathstring/",
					RawQuery: "key=key&token=token",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid URL, unsupported API Version 2",
			args: args{
				u: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/b1/spaces/spaceID/messages/",
					RawQuery: "key=key&token=token",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid URL, unsupported spaces",
			args: args{
				u: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/v1/thisisnotspaces/spaceID/messages/",
					RawQuery: "key=key&token=token",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid URL, SpaceID is empty",
			args: args{
				u: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/v1/spaces/spaceID/wherever/",
					RawQuery: "key=key&token=token",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid URL, SpaceID too long",
			args: args{
				u: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/v1/spaces/asdfjbuasdfnvouhsfdioasdiof/messages/",
					RawQuery: "key=key&token=token",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid URL, messages string is not in the url path",
			args: args{
				u: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/v1/spaces/spaceID/wherever/",
					RawQuery: "key=key&token=token",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid URL, key is not present in the query string",
			args: args{
				u: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/v1/spaces/spaceID/messages/",
					RawQuery: "nokey=key&token=token",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid URL, token is not present in the query string",
			args: args{
				u: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/v1/spaces/spaceID/messages/",
					RawQuery: "key=key&notoken=token",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWebhookURL(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWebhookURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWebhookURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWebhookURL_String(t *testing.T) {
	u := &url.URL{
		Scheme:   "https",
		Host:     "chat.googleapis.com",
		Path:     "/v1/spaces/spaceID/messages/",
		RawQuery: "key=key&token=token",
	}

	got, err := NewWebhookURL(u)
	if err != nil {
		t.Errorf("NewWebhookURL() error = %v", err)
	}

	if got.String() != u.String() {
		t.Errorf("NewWebhookURL.String() = %v, want %v", got, u)
	}
}

func TestNewWebhookURL_GetQuery(t *testing.T) {
	u := &url.URL{
		Scheme:   "https",
		Host:     "chat.googleapis.com",
		Path:     "/v1/spaces/spaceID/messages/",
		RawQuery: "key=key&token=token",
	}

	got, err := NewWebhookURL(u)
	if err != nil {
		t.Errorf("NewWebhookURL() error = %v", err)
	}

	if got.GetQuery() != u.RawQuery {
		t.Errorf("NewWebhookURL.GetQuery() = %v, want %v", got, u)
	}
}

func TestNewWebhookURL_GetSpaceID(t *testing.T) {
	u := &url.URL{
		Scheme:   "https",
		Host:     "chat.googleapis.com",
		Path:     "/v1/spaces/spaceID/messages/",
		RawQuery: "key=key&token=token",
	}

	got, err := NewWebhookURL(u)
	if err != nil {
		t.Errorf("NewWebhookURL() error = %v", err)
	}

	if got.GetSpaceID() != "spaceID" {
		t.Errorf("NewWebhookURL.GetSpaceID() = %v, want %v", got, u)
	}
}

func TestNewWebhookURL_GetKey(t *testing.T) {
	u := &url.URL{
		Scheme:   "https",
		Host:     "chat.googleapis.com",
		Path:     "/v1/spaces/spaceID/messages/",
		RawQuery: "key=key&token=token",
	}

	got, err := NewWebhookURL(u)
	if err != nil {
		t.Errorf("NewWebhookURL() error = %v", err)
	}

	if got.GetKey() != "key" {
		t.Errorf("NewWebhookURL.GetKey() = %v, want %v", got, u)
	}
}

func TestNewWebhookURL_GetToken(t *testing.T) {
	u := &url.URL{
		Scheme:   "https",
		Host:     "chat.googleapis.com",
		Path:     "/v1/spaces/spaceID/messages/",
		RawQuery: "key=key&token=token",
	}

	got, err := NewWebhookURL(u)
	if err != nil {
		t.Errorf("NewWebhookURL() error = %v", err)
	}

	if got.GetToken() != "token" {
		t.Errorf("NewWebhookURL.GetToken() = %v, want %v", got, u)
	}
}
