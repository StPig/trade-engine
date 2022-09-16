# Trade-Engine

## Basic Type
- Order

    To store order properties

    - Name: request name
    - ReqOrdID: request order id
    - Price: price
    - Quantity: quantity
    - OrderType: order type, 0: Market, 1: Limit
    - OrderSide: order side, 0: Buy, 1: Sell
- Trade

    To store trade properties

    - TakerOrdID: taker order id
    - MakerOrdID: maker order id
    - Price: trade price
    - Quantity: trade quantity

## Protocol
- Request
    - CreateOrderRequest
        - name: request name
        - reqOrdID: request order id
        - price: price
        - quantity: quantity
        - orderType: order type, 0: Market, 1: Limit
        - orderSide: order side, 0: Buy, 1: Sell

    example:
    ```
        name: "Test1",
        reqOrdID: 1,
        price: 10,
        quantity: 1,
        orderType: 0,
        orderSide: 0
    ```
- Response
    - ExecutionReport
        - name: request name
        - reqOrdID: request order id
        - status: request status


    example:
    ```
        name: "Test1",
        reqOrdID: 1,
        status: "success"
    ```

## Engine Logic
`Order` includes basic order information, `OrderBook` includes a read write lock and two slices of type `Order`.
`BuyOrders` is sorted by price in ascending order, the highest price is at the end. `SellOrders` is sorted by price in descending order, and the lowest price is at the end.
`addBuyOrder` and `addSellOrder` insert the incoming order into the corresponding position of the slice by price.
`removeBuyOrder` and `removeSellOrder` remove elements based on the input index.

Create a gRPC server and receive requests to create orders via `CreateOrder` function.
In the engine, the `process` function will determine the action is buy or sell, market price or limit price, and direct to the corresponding function to compare orders.
`processMarketBuy` and `processMarketSell` will loop forward from the end of the slice to take the required quantity.
`processLimitBuy` and `processLimitSell` will first confirm whether the last element of the slice matches the price, and then loop forward to take the required quantity until the price no longer matches or the quantity is enough.
When there is a matching order, enter `fillOrder` to calculate the transaction quantity and return `Trade`. If the provided order quantity is zero, call `removeBuyOrder` or `removeSellOrder` to remove the order.
Finally, if there are remaining demand orders, they are added to the queue by `addBuyOrder` and `addSellOrder`.
Then return the transaction result.
