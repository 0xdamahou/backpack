# Backpack API Golang SDK

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.20-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/0xdamahou/backpack)](https://goreportcard.com/report/github.com/0xdamahou/backpack)

## 项目简介

Backpack API Golang SDK 是一个为 [Backpack Exchange](https://backpack.exchange/join/dmh) 提供的非官方 Go 语言客户端库。该项目旨在为 Go 开发者提供一个简洁、高效、类型安全的方式来与 Backpack 交易所进行交互，支持现货交易、账户管理、市场数据获取等全方位功能。

## 核心特性

### 🚀 功能完备
- **REST API 完整覆盖**：支持所有公开和私有 REST 端点
- **WebSocket 实时数据**：提供市场数据和用户数据的实时推送
- **签名认证**：内置 Ed25519 签名机制，确保交易安全
- **错误处理**：完善的错误类型定义和重试机制

### 💡 开发友好
- **类型安全**：充分利用 Go 的强类型特性，减少运行时错误
- **链式调用**：支持流畅的 API 调用体验
- **详细文档**：每个函数都有完整的 GoDoc 注释



## 快速开始

### 安装

```bash
go get github.com/0xdamahou/backpack
```

### 基础使用

```go
package main

import (
   
    "log"
    backpack "github.com/0xdamahou/backpack/authenticated"
)

func main() {
    // 创建客户端
    client,err := backpack.NewBackpackClient( "your-api-key","your-api-secret" )
    if err!=nil{
		log.Println(err)
		return
    }
    // 获取账户余额

    balances, err := client.GetBalances()
    if err != nil {
        log.Fatal(err)
    }
    
    for _, balance := range balances {
        log.Printf("%+v",balance)
    }
}
```

## API 使用示例

### 市场数据

```go
import (
"github/0xdamahou/backpack/public"
)

bpbc := public.NewBackpackPublicClient()
// 获取所有交易对信息
markets, err := bpbc.GetMarkets()

// 获取K线数据
symbol := "SOL_USDC_PERP"
to := time.Now()
from := to.AddDate(-1, -2, -10)
klines, err := bpbc.GetKline(symbol, "1d", from.Unix(), to.Unix())

// 获取订单簿
orderbook, err := client.GetDepth( "SOL_USDC")

```

### 交易功能

```go
// 下单
client,err := backpack.NewBackpackClient( "your-api-key","your-api-secret" )

order, err := client.LimitOrder("SOL_USDC_PERP", OrderSideSell, "3", "144.5")

// 查询订单
order, err := client.GetOpenOrders( "SOL_USDC_PERP")

// 取消订单
cor:=NewOpenOrderRequest("SOL_USDC_PERP")
cor.WithOrderId(orderID)
//或者cor.WithClientId(clientOrderId)
result, err := client.CancelOpenOrder(cor)

// 批量取消订单
results, err := client.CancelAllOrders( "SOL_USDC",backpack.CancelOrderTypeRestingLimit)
```

### WebSocket 订阅

```go
func monitorOrder() {
        client := streams.NewBackpackWebsocketClient(streams.DefaultWsURL, apiKey, apiSecret)
        done := make(chan struct{})
        result := make(chan []byte, 100)
        interrupt := make(chan os.Signal, 1)
        signal.Notify(interrupt, os.Interrupt)
        con, err := client.OrderUpdate("", done, result)
        if err != nil {
        log.Fatal(err)
        return
        }
        defer con.Close()
        for {
                select {
                    case <-interrupt:
                        close(done)
                        return
                
                     case bs, ok := <-result:
                        if !ok {
                            log.Println("Reading from backpack error.")
                            done <- struct{}{}
                            close(done)
                            return
                        } else {
                            var message streams.StreamMessage
                            _ = json.Unmarshal(bs, &message)
                            var order streams.OrderEvent
                            _ = json.Unmarshal(message.Data, &order)
                            
                            //log.Printf("%s %s %s %s %s %s %s\n", order.Symbol, order.Side, order.OrderID, order.Quantity, order.Price, order.Fee, order.FeeSymbol)
                            if order.EventType == "orderFill" {
                            log.Printf("%+v\n", order)
                            }
                        }
                
                    }
            }
}
```





## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 免责声明

本项目是非官方的第三方库，与 Backpack Exchange 无关。使用本库进行交易需自行承担风险。请确保充分了解加密货币交易的风险，并在使用前进行充分的测试。
