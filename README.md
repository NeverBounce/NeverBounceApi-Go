<p align="center"><img src="https://neverbounce-marketing.s3.amazonaws.com/neverbounce_color_600px.png"></p>

<p align="center">
<a href="https://travis-ci.org/NeverBounce/NeverBounceApi-Go"><img src="https://travis-ci.org/NeverBounce/NeverBounceApi-Go.svg" alt="Build Status"></a>
<a href="https://codeclimate.com/github/NeverBounce/NeverBounceApi-Go"><img src="https://codeclimate.com/github/NeverBounce/NeverBounceApi-Go/badges/gpa.svg" /></a>
</p>

> This version of the wrapper is for the V4 API currently in beta

# neverbounce
--
    import "github.com/NeverBounce/NeverBounceApi-Go"

Package neverbounce creates native Golang mappings to use NeverBounce's email
verification API. Our verification API allows you to create Custom Integrations
to add email verification to any part of your software. We offer solutions for
verifying individual emails as well as lists containing hundreds or even
millions of emails.

For our full API documentation see: https://developers.neverbounce.com/v4.0/

Basic usage:

    import "github.com/neverbounce/neverbounceapi-go"
    client, err := neverbounce.New("api_key")
    if err != nil {
    	panic(err)
    }

    accountInfo, err := client.Account.Info()
    if err != nil {
    	panic(err)
    }
    fmt.Println(accountInfo)


Additional examples can be found in the examples directory

## Usage

```go
const DefaultBaseURL = "https://api.neverbounce.com/v4/"
```
DefaultBaseURL is the default host to make the API requests on

#### func  MakeRequest

```go
func MakeRequest(method string, url string, data interface{}) ([]byte, error)
```
MakeRequest handles the request and parsing of the responses to and from the API
It will throw and error when a 4xx/5xx HTTP code is encountered or if the API
returns an api error. See:
https://developers.neverbounce.com/v4.0/reference#error-handling

### type Account

```go
type Account struct {
}
```

Account contains bindings for account related API endpoints.

#### func (*Account) Info

```go
func (r *Account) Info() (*nbModels.AccountInfoResponseModel, error)
```
Info returns the account's current credit balance as well as job counts
indicating the number of jobs currently in the account.

### type Jobs

```go
type Jobs struct {
}
```

Jobs endpoints provides high-speed​ validation on a list of email addresses. You
can use the status endpoint to retrieve real-time statistics about a bulk job in
progress. Once the job has finished, the results can be retrieved with our
download endpoint.

#### func (*Jobs) CreateFromRemoteURL

```go
func (r *Jobs) CreateFromRemoteURL(model *nbModels.JobsCreateRemoteURLRequestModel) (*nbModels.JobsCreateResponseModel, error)
```
CreateFromRemoteURL creates a new job from a comma separated value (CSV) file
hosted on a remote URL. The URL supplied can be any commonly available protocal;
e.g: HTTP, HTTPS, FTP, SFTP. Basic auth is supported by including the
credentials in the URI string; e.g:
http://name:passwd@example.com/full/path/to/file.csv

#### func (*Jobs) CreateFromSuppliedData

```go
func (r *Jobs) CreateFromSuppliedData(model *nbModels.JobsCreateSuppliedDataRequestModel) (*nbModels.JobsCreateResponseModel, error)
```
CreateFromSuppliedData creates a new job from data you supply directly in the
request. Supplied data will need to be given as a map, see the examples in the
nbModel package.

#### func (*Jobs) Delete

```go
func (r *Jobs) Delete(model *nbModels.JobsDeleteRequestModel) (*nbModels.JobsDeleteResponseModel, error)
```
Delete will remove the job and it's verification data (if previously verified)
This can only be done when a job is Queued, Waiting, Completed, or Failed. A job
cannot be deleted while it is being uploaded, parsed, or ran. Once deleted the
job results cannot be recovered, deletion is permanent.

#### func (*Jobs) Download

```go
func (r *Jobs) Download(model *nbModels.JobsDownloadRequestModel, filepath string) error
```
Download the results as a CSV to a file. This is useful if your uploading the
results to a CRM or are use to working with spreadsheets.

#### func (*Jobs) Parse

```go
func (r *Jobs) Parse(model *nbModels.JobsParseRequestModel) (*nbModels.JobsParseResponseModel, error)
```
Parse allows you to parse the job data after creation. If you create a job with
AutoParse set to true (defaults to false) you do not need to use this method.
Once parsed, a job cannot be reparsed.

#### func (*Jobs) Results

```go
func (r *Jobs) Results(model *nbModels.JobsResultsRequestModel) (*nbModels.JobsResultsResponseModel, error)
```
Results will return the actual verification results. This can only be done once
the job has reached the completed status. The results will be returned in
batches according to the pagination options you've supplied. Verification info
will be formatted the same way Single.Check returns verification info.

#### func (*Jobs) Search

```go
func (r *Jobs) Search(model *nbModels.JobsSearchRequestModel) (*nbModels.JobsSearchResponseModel, error)
```
Search the jobs you've previously submitted to your account. It will return jobs
in batches according to the pagination options you've supplied. The returned
jobs will include the same information available from the Status method

#### func (*Jobs) Start

```go
func (r *Jobs) Start(model *nbModels.JobsStartRequestModel) (*nbModels.JobsStartResponseModel, error)
```
Start allows you to start a job after it has been parsed. If you create a job or
parse a job with AutoStart set to true (defaults to false) you do not need to
use this method. Once the list has been started the credits will be deducted and
the process cannot be stopped or restarted.

#### func (*Jobs) Status

```go
func (r *Jobs) Status(model *nbModels.JobsStatusRequestModel) (*nbModels.JobsStatusResponseModel, error)
```
Status will return information pertaining to the Jobs state. It will include the
jobs current status as well as the verification stats. This will be the primary
property you'll want to check to determine what can be done with the job.

### type NeverBounce

```go
type NeverBounce struct {
	Account *Account
	Single  *Single
	Jobs    *Jobs
	POE     *POE
}
```

NeverBounce is the root struct of the wrapper. This is used to access the
specific bindings.

#### func  New

```go
func New(apiKey string) (*NeverBounce, error)
```
New creates a new instance of *NeverBounce. Accepts the api key to use for
authentication.

#### func (*NeverBounce) SetBaseURL

```go
func (r *NeverBounce) SetBaseURL(url string)
```
SetBaseURL will set the url used to make the requests (overrides the
DefaultBaseURL constant). This method is primarily for internal testing and
debugging purposes, under normal circumstances it will not be used

### type POE

```go
type POE struct {
}
```

POE endpoints allow you to confirm frontend verifications performed by the
Javascript Widget

#### func (*POE) Confirm

```go
func (r *POE) Confirm(model *nbModels.POEConfirmRequestModel) (*nbModels.POEConfirmResponseModel, error)
```
Confirm verifies that the result provided during frontend verification (e.g.
Javascript Widget) has not been tampered with. It requires you to pass the
email, result, transaction_id, and confirmation_token supplied by the
verification.

### type Single

```go
type Single struct {
}
```

Single endpoints allow you to integrate our email verification into your
existing applications at the point of entry and onboarding processes

#### func (*Single) Check

```go
func (r *Single) Check(model *nbModels.SingleCheckRequestModel) (*nbModels.SingleCheckResponseModel, error)
```
Check verifies the email provided and returns the verification result. In
addition to this, it can also return a breakdown of the email address' host info
and your account balance
