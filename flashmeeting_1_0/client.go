// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package flashmeeting_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateFlashMeetingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateFlashMeetingHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateFlashMeetingHeaders) GoString() string {
	return s.String()
}

func (s *CreateFlashMeetingHeaders) SetCommonHeaders(v map[string]*string) *CreateFlashMeetingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateFlashMeetingHeaders) SetXAcsDingtalkAccessToken(v string) *CreateFlashMeetingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateFlashMeetingRequest struct {
	// 创建人union id
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 日程id
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// 钉闪会名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateFlashMeetingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlashMeetingRequest) GoString() string {
	return s.String()
}

func (s *CreateFlashMeetingRequest) SetCreator(v string) *CreateFlashMeetingRequest {
	s.Creator = &v
	return s
}

func (s *CreateFlashMeetingRequest) SetEventId(v string) *CreateFlashMeetingRequest {
	s.EventId = &v
	return s
}

func (s *CreateFlashMeetingRequest) SetTitle(v string) *CreateFlashMeetingRequest {
	s.Title = &v
	return s
}

type CreateFlashMeetingResponseBody struct {
	// 闪会结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 闪会的key
	FlashMeetingKey *string `json:"flashMeetingKey,omitempty" xml:"flashMeetingKey,omitempty"`
	// 闪会开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 闪会标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 闪会url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateFlashMeetingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlashMeetingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlashMeetingResponseBody) SetEndTime(v int64) *CreateFlashMeetingResponseBody {
	s.EndTime = &v
	return s
}

func (s *CreateFlashMeetingResponseBody) SetFlashMeetingKey(v string) *CreateFlashMeetingResponseBody {
	s.FlashMeetingKey = &v
	return s
}

func (s *CreateFlashMeetingResponseBody) SetStartTime(v int64) *CreateFlashMeetingResponseBody {
	s.StartTime = &v
	return s
}

func (s *CreateFlashMeetingResponseBody) SetTitle(v string) *CreateFlashMeetingResponseBody {
	s.Title = &v
	return s
}

func (s *CreateFlashMeetingResponseBody) SetUrl(v string) *CreateFlashMeetingResponseBody {
	s.Url = &v
	return s
}

type CreateFlashMeetingResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlashMeetingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlashMeetingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlashMeetingResponse) GoString() string {
	return s.String()
}

func (s *CreateFlashMeetingResponse) SetHeaders(v map[string]*string) *CreateFlashMeetingResponse {
	s.Headers = v
	return s
}

func (s *CreateFlashMeetingResponse) SetBody(v *CreateFlashMeetingResponseBody) *CreateFlashMeetingResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

func (client *Client) CreateFlashMeeting(request *CreateFlashMeetingRequest) (_result *CreateFlashMeetingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateFlashMeetingHeaders{}
	_result = &CreateFlashMeetingResponse{}
	_body, _err := client.CreateFlashMeetingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlashMeetingWithOptions(request *CreateFlashMeetingRequest, headers *CreateFlashMeetingHeaders, runtime *util.RuntimeOptions) (_result *CreateFlashMeetingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		body["creator"] = request.Creator
	}

	if !tea.BoolValue(util.IsUnset(request.EventId)) {
		body["eventId"] = request.EventId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateFlashMeetingResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateFlashMeeting"), tea.String("flashmeeting_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/flashmeeting/meetings"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
