Feature: Find an existing student
  In order to find an existing student
  As a administrator
  I want to use a GET request to '/students/:id'

  Scenario: Find an existing student
    Given I send a GET request to '/students/45215570-0296-11ee-8566-acde48001122'
    Then the response status should be 200