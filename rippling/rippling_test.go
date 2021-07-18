package rippling

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/cucumber/godog"
	"github.com/jarcoal/httpmock"
)

type response struct {
	payload string
	status int
}
type request struct {
	token string
	url string
	report string
	reportDate string
}
type ripplingTest struct {
	client *Client
	resp response
	req request	
}

func (r *ripplingTest) s(arg1 string) error {
	return godog.ErrPending
}

func(r *ripplingTest) i() error {

		urlRegexp, _ := regexp.Compile(fmt.Sprintf(`^%s`, r.req.url))
		
		httpmock.RegisterRegexpResponder("GET", urlRegexp,
			httpmock.NewStringResponder(200, `{"foo":"bar"}`))
		// r.Client.GetData(r.req.token, false)
		callCount := httpmock.GetTotalCallCount()
		if callCount != 1 {
			return godog.ErrUndefined	
		}
		return godog.ErrPending
}

func (r *ripplingTest) b(arg1 string) error {
	return godog.ErrPending
}

func (r *ripplingTest) c(arg1, arg2, arg3, arg4, arg5, arg6 string) error {
	return godog.ErrPending
}

func (r *ripplingTest) d(url, report, reportDate, token, statusStr, payload string) error {
	r.req = request{
		token: token,
		url: url,
		reportDate: reportDate,
		report: report,
	}
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		status = 200
	}
	r.resp = response{
		payload: payload,
		status: status,
	}
	return godog.ErrPending
}

func Z(ctx *godog.ScenarioContext) {
	// rip := ripplingTest{}
	// fmt.Print(rip)
	// httpmock.Activate()
	// defer httpmock.DeactivateAndReset()
	
	// ctx.Step(`^I am fetching the report "([^"]*)" for "([^"]*)" with the token "([^"]*)"$`, rip.iAmFetchingTheReportForWithTheToken)
	// ctx.Step(`^I call GetData$`, rip.iCallGetData)
	
	// ctx.Step(`^the err response should be "([^"]*)"$`, rip.theErrResponseShouldBe)
	// ctx.Step(`^the err response should be ""([^"]*)""$`, rip.theErrResponseShouldBe)
	// ctx.Step(`^the response should equal "([^"]*)"$`, rip.theResponseShouldEqual)
	// ctx.Step(`^the response should equal "\[{"([^"]*)": "([^"]*)","comp":"([^"]*)","type":"([^"]*)","workStart":"([^"]*)","workEnd":"([^"]*)"}\]"$`, rip.theResponseShouldEqualCompTypeWorkStartWorkEnd)
	// ctx.Step(`^we are calling the URL "([^"]*)" the report "([^"]*)" on "([^"]*)" the token "([^"]*)", and the API responds with the status "([^"]*)" and the payload "([^"]*)"$`, rip.weAreCallingTheURLTheReportOnTheTokenAndTheAPIRespondsWithTheStatusAndThePayload)
}