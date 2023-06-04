Feature: Find an existing student
  In order to find an existing student
  As a administrator
  I want to use a GET request to '/students/:id'

  Scenario: Find an existing student
    Given I send a GET request to '/students/fe7017d8-9e8f-4952-e047e36b1694'
    Then the response status should be 200
    And the response body should be a JSON object
    And the response body should have a student with the following attributes:
      """
       {
          'student_id': 'fe7017d8-9e8f-4952-e047e36b1694',
          'student_name': 'John Doe',
          'student_age": 20
        }
      """