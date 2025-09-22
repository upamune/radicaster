# RSS Feed Filtering

This feature allows filtering RSS feeds by time period using the `since` query parameter.

## Usage

Add the `since` query parameter to any RSS feed URL to filter episodes:

### Examples

- Get episodes from the last year: `/rss.xml?since=1y`
- Get episodes from the last 6 months: `/rss.xml?since=6m`
- Get episodes from the last 30 days: `/rss.xml?since=30d`
- Get episodes from the last 24 hours: `/rss.xml?since=24h`

### Supported Time Units

- `h` - hours (e.g., `24h`)
- `d` - days (e.g., `30d`)
- `m` - months (e.g., `6m`, approximate as 30 days)
- `y` - years (e.g., `1y`, approximate as 365 days)

### Endpoints that support filtering

- `/rss.xml?since=1y` - Default RSS feed filtered to last year
- `/:program_path/rss.xml?since=6m` - Program-specific feed filtered to last 6 months
- `/zenroku/:program_path/rss.xml?since=30d` - Zenroku program feed filtered to last 30 days

### Backward Compatibility

When no `since` parameter is provided, the RSS feed will contain all episodes as before.

If an invalid `since` parameter is provided, the system will log a warning and return the unfiltered feed.