# Rss Feed Agregator API

This is a rest api that enables users to creat accounts and create rss feeds in their accounts.
It also enables users to retrieve and delete their rss feeds. 

Users can also get the posts associated with their rss feeds and follows other users rss feeds.

**Note: Only users with an account can create and delete rss feeds as well as follow other rss feeds**

## API Documentation

This api is currently not live. To use it, begin by downloading the repository in your local machine.
After downloading, create a **.env file** in the root directory and add a **PORT** 
and **DB_URL** variable for the API to listen on and the postgress database configuration string
for it to connect to your postgres database 

### Example

```
    PORT= 8000
    DB_URL= postgres://your_postgres_username:your_postgres_password@localhost:5432/**rssagg**?sslmode=disable
```

**Notice the database name rssagg in the connection string must be exactly named rssagg since that's what's configured**
**in the code unless you decide to change the name in internal/databade/config file**

### How to run the aplication

    ```shell
        go build && ./rssagg.exe (windows)
        go build && ./rssagg (mac os x)
    ```

### 1 - Create a user

```JSON
    URL - http://localhost:8000/v1/users
    Method - POST
    Headers-
        {
            content-type: application/json
        }
    Body - 
        {
            "name": "Miles Morales"
        }
```

**Example Response Body**
```JSON
    {
        "id": "542656d6-92df-47d8-b6fc-f18281aecf8f",
        "created_at": "2023-06-21T07:20:38.97302Z",
        "updated_at": "2023-06-21T07:20:38.97302Z",
        "name": "Miles Morales",
        "api_key": "936d3618d3d171ee65fdf53605693e43c113c67dc249a95afe9f3773b1dcacff"
    }
```

**Please, ensure you save your api key as that is required for all authenticated endpoints.**

### 2 - Retrieve a user

Only the owner can retrieve and see user information. Therefore, this is an authenticated endpoint and requires
the api key 

```JSON
    URL - http://localhost:8000/v1/users
    Method - GET
    Headers-
        {
            content-type: application/json,
            Authorization:  Apikey 936d3618d3d171ee65fdf53605693e43c113c67dc249a95afe9f3773b1dcacff
        }
```

**Example Response Body**

```JSON
    {
        "id": "fd907363-e4dd-4868-9361-41145ed5c107",
        "created_at": "2023-06-21T08:30:16.942314Z",
        "updated_at": "2023-06-21T08:30:16.942314Z",
        "name": "Miles Morales",
        "api_key": "6dc7b055c14d5503640f6d6852e63065f76704bdeb57075bb60a2d1cd1cf5683"
    }
```

**Please, ensure you write the value for the Authorization header in the format as in the above example, otherwise you** 
**will get a 400 error. This is true for all authenticated endpoints.**

### 3 - Create a Feed

    
```JSON
    URL- http://localhost:8000/v1/feeds
    Method - POST
    Headers - 
        {
            content-type: application/json,
            Authorization:  Apikey 936d3618d3d171ee65fdf53605693e43c113c67dc249a95afe9f3773b1dcacff
        }
    Body -
        {
            "name": "Nasa Education News Feed",
            "url": "https://www.nasa.gov/rss/dyn/educationnews.rss"
        }
```

**Example Response Body**  
```JSON
    {
        "id": "084f08b0-03ed-4bb1-9e3c-5127e299c199",
        "created_at": "2023-06-21T08:20:21.474941Z",
        "updated_at": "2023-06-21T08:20:21.474941Z",
        "name": "Nasa Education News Feed",
        "url": "https://www.nasa.gov/rss/dyn/educationnews.rss",
        "user_id": "542656d6-92df-47d8-b6fc-f18281aecf8f",
        "last_fetched_at": "2023-06-21T08:20:21.474941Z"
    }
```

**Since feeds are personal to a user, meaning the url of a feed is unique, you cannot create feeds with urls that**
**already exist in the dabase. Doing so, will result in the following response:**

```JSON

    {
        "error": "Couldn't create feed. url is already associated with an existing feed."
    }
```

### 4 - Retrieve all feeds

This is not an authenticated endpoint. Making a request to this endpoint will retrieve all feeds in the database, including those of other users.

```JSON
    URL- http://localhost:8000/v1/feeds
    Method - GET
    Headers - 
        {
            content-type: application/json,
        }

```
**Example Response Body** 

```JSON
    [
        {
            "id": "c224fc5a-6506-4d5b-a67d-b20f9aa9de02",
            "created_at": "2023-05-22T23:55:42.908699Z",
            "updated_at": "2023-06-21T01:21:16.587176Z",
            "name": "Lane's Blog",
            "url": "https://wagslane.dev/index.xml",
            "user_id": "b1e0944e-2f9c-4570-8e59-d8fdc9af8b5b",
            "last_fetched_at": "2023-06-21T01:21:16.587176Z"
        },
        {
            "id": "084f08b0-03ed-4bb1-9e3c-5127e299c199",
            "created_at": "2023-06-21T08:20:21.474941Z",
            "updated_at": "2023-06-21T01:21:16.587178Z",
            "name": "Nasa Education News Feed",
            "url": "https://www.nasa.gov/rss/dyn/educationnews.rss",
            "user_id": "542656d6-92df-47d8-b6fc-f18281aecf8f",
            "last_fetched_at": "2023-06-21T01:21:16.587178Z"
        },
        {
            "id": "566fd6cd-ebc4-496a-a169-96f5fba777b6",
            "created_at": "2023-05-23T00:03:14.708708Z",
            "updated_at": "2023-06-21T01:21:16.606989Z",
            "name": "Nasa Breaking News Feed",
            "url": "https://www.nasa.gov/rss/dyn/breaking_news.rss",
            "user_id": "cd834ea2-ac52-4779-a380-bfdce6b7b41b",
            "last_fetched_at": "2023-06-21T01:21:16.606989Z"
        }
    ]
```

### 5 - Follow a feed 

Users can follow other users feed

```JSON
    URL- http://localhost:8000/v1/feed_follows
    Method - POST
    Headers - 
        {
            content-type: application/json,
            Authorization:  Apikey 936d3618d3d171ee65fdf53605693e43c113c67dc249a95afe9f3773b1dcacff
        }
    Body -
        {
            "feed_id": "566fd6cd-ebc4-496a-a169-96f5fba777b6"
        }
```

**Example Response Body** 

```JSON
    {
        "msg": "You are following the given feed now",
        "following": {
            "id": "4be9a3e2-29a1-4d12-b797-81a6fffddd4e",
            "created_at": "2023-06-21T02:39:49.5160014-07:00",
            "updated_at": "2023-06-21T02:39:49.5160014-07:00",
            "user_id": "542656d6-92df-47d8-b6fc-f18281aecf8f",
            "feed_id": "566fd6cd-ebc4-496a-a169-96f5fba777b6"
        }
    }
```

### 6 - Retrieve all feeds that a user is following

```JSON
    URL- http://localhost:8000/v1/feed_follows
    Method - GET
    Headers - 
        {
            content-type: application/json,
            Authorization:  Apikey 936d3618d3d171ee65fdf53605693e43c113c67dc249a95afe9f3773b1dcacff
        }
```

**Example Response Body** 

```JSON
    {
        "following": [
            {
                "id": "5c4c6476-b091-48e3-bca4-599bd4ed8f2e",
                "created_at": "2023-05-23T22:55:38.776482Z",
                "updated_at": "2023-05-23T22:55:38.776482Z",
                "user_id": "96e93272-c5ff-465d-8141-1bef8d737339",
                    "feed_id": "c224fc5a-6506-4d5b-a67d-b20f9aa9de02"
            },
            {
                "id": "fd34c68a-3517-4c58-bd43-dc8fa2c2f32a",
                "created_at": "2023-05-24T01:02:39.700368Z",
                "updated_at": "2023-05-24T01:02:39.700368Z",
                "user_id": "96e93272-c5ff-465d-8141-1bef8d737339",
                "feed_id": "566fd6cd-ebc4-496a-a169-96f5fba777b6"
            }
        ]
    }
```

### 7 - Retrieve all posts from feeds that a user is following
    
Since every rss feed contains one or more posts, a user can retrieve all posts from all rss feeds they are following.

```JSON
    URL- http://localhost:8000/v1/posts
    Method - GET
    Headers - 
        {
            content-type: application/json,
            Authorization:  Apikey 936d3618d3d171ee65fdf53605693e43c113c67dc249a95afe9f3773b1dcacff
        }
```

**Example Response Body** 

```JSON
    [
        {
            "id": "2b99c191-3117-448c-a64f-ed392239e24a",
            "created_at": "2023-06-16T10:03:40.250385Z",
            "updated_at": "2023-06-16T10:03:40.250385Z",
            "title": "The Zen of Proverbs",
            "description": "20 rules of thumb for writing better software.\n Optimize for simplicity first Write code for humans, not computers Reading is more important than writing Any style is fine, as long as it&rsquo;s black There should be one way to do it, but seriously this time Hide the sharp knives Changing the rules is better than adding exceptions Libraries are better than frameworks Transitive dependencies are a problem Dynamic runtime dependencies are a bigger problem API surface area is a liability Returning early is a good thing Use more plain text Compiler errors are better than runtime errors Runtime errors are better than bugs Tooling is better than documentation Documentation is better than nothing Configuration sucks, but so does convention The cost of building a feature is its smallest cost Types are one honking great idea &ndash; let&rsquo;s do more of those!",
            "published_at": "2023-01-08T00:00:00Z",
            "url": "https://wagslane.dev/posts/zen-of-proverbs/",
            "feed_id": "c224fc5a-6506-4d5b-a67d-b20f9aa9de02"
        },
        {
            "id": "65d6eca3-afc9-430a-9986-95ac51f86092",
            "created_at": "2023-06-16T10:03:40.262337Z",
            "updated_at": "2023-06-16T10:03:40.262337Z",
            "title": "College: A Solution in Search of a Problem",
            "description": "College has been prescribed almost universally by the parents of the last ~40 years as the solution to life&rsquo;s problems. We&rsquo;ve been told it&rsquo;s the way to land a good job and to make more money. But is it?\nI think that most college degrees these days commit a cardinal sin in the business world. College degrees are solutions in search of a problem.\nWhat is a solution in search of a problem?",
            "published_at": "2022-12-17T00:00:00Z",
            "url": "https://wagslane.dev/posts/college-a-solution-in-search-of-a-problem/",
            "feed_id": "c224fc5a-6506-4d5b-a67d-b20f9aa9de02"
        },
        {
            "id": "f1fb11e3-7eca-4609-b68b-2240288631b8",
            "created_at": "2023-06-16T10:03:40.26284Z",
            "updated_at": "2023-06-16T10:03:40.26284Z",
            "title": "Thoughts on the \"Guard\" Proposal for Go's Error Handling",
            "description": "I found this proposal for improvements to error handling in Go interesting, but still not something I&rsquo;d be happy to see implemented.\nAllow me to clear up my thoughts on Go&rsquo;s errors. Overall, I prefer how Go forces me to think about errors at every turn. When working in try/catch languages like JavaScript, I often easily forget which functions can throw. Even if I do remember, it&rsquo;s easy to think &ldquo;I think this gets caught somewhere up the call chain&rdquo;.",
            "published_at": "2022-11-05T00:00:00Z",
            "url": "https://wagslane.dev/posts/guard-keyword-error-handling-golang/",
            "feed_id": "c224fc5a-6506-4d5b-a67d-b20f9aa9de02"
        },
        .
        .
        .
        {
            "id": "ab4a0dcb-53b1-43bc-90ef-a3702b45939f",
            "created_at": "2023-06-16T10:03:40.26284Z",
            "updated_at": "2023-06-16T10:03:40.26284Z",
            "title": "Learn to Say 'No'",
            "description": "Saying &ldquo;no&rdquo; is a hard skill to learn. It&rsquo;s even harder if you tend to be a more introverted person. However, learning how to say &ldquo;no&rdquo; effectively can help your career. I certainly have struggled over the years with saying &ldquo;no&rdquo; as a programmer, after all, wouldn&rsquo;t a good programmertm be able to do anything?\nLet&rsquo;s look at some example scenarios where perhaps you should be saying &ldquo;no&rdquo; more often.",
            "published_at": "2022-06-27T00:00:00Z",
            "url": "https://wagslane.dev/posts/developers-learn-to-say-no/",
            "feed_id": "c224fc5a-6506-4d5b-a67d-b20f9aa9de02"
        },
    ]
```