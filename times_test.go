package humanize

import (
	"testing"
	"time"
)

type timeTest struct {
	now time.Time
	t   time.Time
	s   string
}

var now = time.Now()

var timeTests = []timeTest{
	{now, now, "now"},
	{now, now.Add(-1 * time.Second), "1 second ago"},
	{now, now.Add(-12 * time.Second), "12 seconds ago"},
	{now, now.Add(-30 * time.Second), "30 seconds ago"},
	{now, now.Add(-45 * time.Second), "45 seconds ago"},
	{now, now.Add(-63 * time.Second), "1 minute ago"},
	{now, now.Add(-15 * time.Minute), "15 minutes ago"},
	{now, now.Add(-63 * time.Minute), "1 hour ago"},
	{now, now.Add(-2 * time.Hour), "2 hours ago"},
	{now, now.Add(-21 * time.Hour), "21 hours ago"},
	{now, now.Add(-26 * time.Hour), "1 day ago"},
	{now, now.Add(-49 * time.Hour), "2 days ago"},
	{now, now.Add(-3 * Day), "3 days ago"},
	{now, now.Add(-7 * Day), "1 week ago"},
	{now, now.Add(-12 * Day), "1 week ago"},
	{now, now.Add(-15 * Day), "2 weeks ago"},
	{now, now.Add(-39 * Day), "1 month ago"},
	{now, now.Add(-99 * Day), "3 months ago"},
	{now, now.Add(-365 * Day), "1 year ago"},
	{now, now.Add(-400 * Day), "1 year ago"},
	{now, now.Add(-548 * Day), "2014-06-12"},
	{now, now.Add(-725 * Day), "2013-12-17"},
	{now, now.Add(-800 * Day), "2013-10-03"},
	{now, now.Add(-3 * Year), "2012-12-27"},
	{now, now.Add(1 * time.Second), "1 second from now"},
	{now, now.Add(12 * time.Second), "12 seconds from now"},
	{now, now.Add(30 * time.Second), "30 seconds from now"},
	{now, now.Add(45 * time.Second), "45 seconds from now"},
	{now, now.Add(63 * time.Second), "1 minute from now"},
	{now, now.Add(15 * time.Minute), "15 minutes from now"},
	{now, now.Add(63 * time.Minute), "1 hour from now"},
	{now, now.Add(2 * time.Hour), "2 hours from now"},
	{now, now.Add(21 * time.Hour), "21 hours from now"},
	{now, now.Add(26 * time.Hour), "1 day from now"},
	{now, now.Add(49 * time.Hour), "2 days from now"},
	{now, now.Add(3 * Day), "3 days from now"},
	{now, now.Add(7 * Day), "1 week from now"},
	{now, now.Add(12 * Day), "1 week from now"},
	{now, now.Add(15 * Day), "2 weeks from now"},
	{now, now.Add(39 * Day), "1 month from now"},
	{now, now.Add(99 * Day), "3 months from now"},
	{now, now.Add(365 * Day), "1 year from now"},
	{now, now.Add(400 * Day), "1 year from now"},
	{now, now.Add(548 * Day), "2017-06-12"},
	{now, now.Add(725 * Day), "2017-12-06"},
	{now, now.Add(800 * Day), "2018-02-19"},
	{now, now.Add(3 * Year), "2018-11-26"},
}

func testTime(t *testing.T, te *timeTest) {
	s := TimeRelativeTo(te.now, te.t)
	if s != te.s {
		t.Errorf("on %s to %s: expected %q but got %q\n",
			te.now.Format(time.RFC3339), te.t.Format(time.RFC3339),
			te.s, s)
	}
}

func TestTime(t *testing.T) {
	for _, te := range timeTests {
		testTime(t, &te)
	}
}
