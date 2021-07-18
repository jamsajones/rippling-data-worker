Feature: Rippling API
  Given the base URL is https://sandbox.rippling.com/api/platform/api
  And the API responds with:
    | ForURL                                   | Method | Token | Report        | ReportDate               | ApiResponse                                                                                        | Status | GetDataResponse                                                                                                            | GetDataError          |
    | /          | GET | 123456 | Hourly_Pay_FC | 2021-05-01T00:00:00.000Z | name,comp,type,workStart,workEnd\njames,4,hourly,2021-05-01T00:00:00.000Z,2021-05-01T00:00:00.000Z | 200    | [{"name": "james","comp":"4","type":"hourly","workStart":"2021-05-01T00:00:00.000Z","workEnd":"2021-05-01T00:00:00.000Z"}] |                       |
    | / | 456789 | Post | Hourly_Pay_CS | 2021-05-01T00:00:00.000Z | 403: Unauthorized Access                                                                           | 403    |                                                                                                                            | "Unauthroized Access" |


  Scenario Outline: calling with token
    Given I am fetching use the token "<Token>"
    When  I call GetData "<Report>" for "<ReportDate>"
    Then  the response should equal "<GetDataResponse>"
      And the err response should be "<GetDataError>"

    Examples:
      | Token  | Report        | ReportDate               | GetDataResponse                                                                                                            | GetDataError          |
      | 123456 | Hourly_Pay_FC | 2021-05-01T00:00:00.000Z | [{"name": "james","comp":"4","type":"hourly","workStart":"2021-05-01T00:00:00.000Z","workEnd":"2021-05-01T00:00:00.000Z"}] |                       |
      | 456789 | Hourly_Pay_CS | 2021-05-01T00:00:00.000Z |                                                                                                                            | "Unauthroized Access" |
