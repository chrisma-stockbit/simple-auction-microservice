# Simple Auction Microservice

I am using this repo to create simple auction microservices in purpose to help myself learn gRPC, go-kit and Go languange. I use MySQL as the datasore.

There will be two microservices, Bidder service and Auction service.\
Bidder service with three endpoints:
- register a bidder (simple RPC)
- top up balance of a bidder (simple RPC)
- get list of successful biddings of a bidder (server side streaming RPC)

Auction Service with three endpoints:
- get list of running auctions (server side streaming RPC)
- initiate auction of an item (simple RPC)
- request to join multiple auctions (client side streaming RPC)
- bid on auction item (bidirectional streaming RPC)

### Server
The server will run on localhost on port 10000.

### Data Model
<pre>
Bidder{
  id string
  full_name string
  account_number string
  bank_name string
  phone_number string
  balance uint
}

Item{
  id string
  name string
  production_year uint
}

Auction {
  id string
  item_id string
  started_at timestamp
  period_in_second uint
  is_running boolean
  final_bid_id string
}

AuctionBidding {
  id string
  auction_id string
  bidder_id string
  price uint
}
</pre>
