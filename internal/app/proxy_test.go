package app

import (
	"net/http"
	"testing"
)

func TestConstraintMatches(t *testing.T) {
	given, when, then, teardown := NewProxyStage(t)
	defer teardown()

	given.a_pact_that_allows_any_names()

	when.
		a_constaint_is_added("sam").and().
		a_request_is_sent_using_the_name("sam")

	then.pact_verification_is_successful()
}

func TestConstraintDoesntMatch(t *testing.T) {
	given, when, then, teardown := NewProxyStage(t)
	defer teardown()

	given.a_pact_that_allows_any_names()

	when.
		a_constaint_is_added("sam").and().
		a_request_is_sent_using_the_name("bob")

	then.pact_verification_is_not_successful()
}

func TestWaitForPact(t *testing.T) {
	given, when, then, teardown := NewProxyStage(t)
	defer teardown()

	given.a_pact_that_allows_any_names()

	when.multiple_requests_are_sent(50)

	then.the_proxy_waits_for_all_requests()
}

func TestWaitForAllPacts(t *testing.T) {
	given, when, then, teardown := NewProxyStage(t)
	defer teardown()

	given.
		a_pact_that_allows_any_names().and().
		a_pact_that_allows_any_address()

	when.requests_for_names_and_addresse_are_sent()

	then.the_proxy_waits_for_all_requests()
}

func TestModifiedStatusCode(t *testing.T) {
	given, when, then, teardown := NewProxyStage(t)
	defer teardown()

	given.
		a_pact_that_allows_any_names().and().
		a_modified_response_status_of_(http.StatusInternalServerError)

	when.a_request_is_sent_with_modifiers_using_the_name("sam")

	then.the_response_is_(http.StatusInternalServerError).and().
		pact_verification_is_successful()

}

func TestModifiedStatusCode_ForNRequests(t *testing.T) {
	given, when, then, teardown := NewProxyStage(t)
	defer teardown()

	given.
		a_pact_that_allows_any_names().and().
		a_modified_response_status_of_(http.StatusInternalServerError).and().
		a_modified_response_attempt_of(2)

	when.n_requests_are_sent_with_modifiers_using_the_name(3, "sam")

	then.
		n_responses_were_received(3).and().
		the_nth_response_is_(1, http.StatusOK).and().
		the_nth_response_is_(2, http.StatusInternalServerError).and().
		the_nth_response_is_(3, http.StatusOK).and().
		pact_verification_is_successful()
}
