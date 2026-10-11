<div align="center">

# go-bitkub

**API ของ Bitkub ครบทุกตัวใน Go module เดียว พร้อม CLI ที่ใช้งานได้ใน 30 วินาที**

ดูราคาคริปโตเป็นเงินบาท สตรีม ticker แบบ real-time และจัดการบัญชี Bitkub ได้ทั้งจากโค้ด Go และ terminal ไม่ต้องเขียน HMAC signing เอง ไม่ต้องแก้บั๊กเวลาเซิร์ฟเวอร์เพี้ยน และไม่ต้องมี API key เพื่อดูข้อมูลตลาด

[![Stars](https://img.shields.io/github/stars/dvgamerr-app/go-bitkub?style=for-the-badge&color=00a86b)](https://github.com/dvgamerr-app/go-bitkub/stargazers)
[![Release](https://img.shields.io/github/v/release/dvgamerr-app/go-bitkub?style=for-the-badge&color=00a86b)](https://github.com/dvgamerr-app/go-bitkub/releases/latest)
[![Go Reference](https://img.shields.io/badge/go-reference-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://pkg.go.dev/github.com/dvgamerr-app/go-bitkub)
[![Licence](https://img.shields.io/github/license/dvgamerr-app/go-bitkub?style=for-the-badge&color=00a86b)](LICENSE)
[![Platforms](https://img.shields.io/badge/linux%20%7C%20macOS%20%7C%20windows-amd64%20%7C%20arm64-555?style=for-the-badge)](https://github.com/dvgamerr-app/go-bitkub/releases/latest)

[เว็บไซต์](https://dvgamerr.app/go-bitkub/) · [เริ่มใช้งาน](#เริ่มใช้งาน) · [คู่มือติดตั้ง](INSTALL.md) · [คู่มือ CLI](CLI.md) · [English](README.md)

<img src="docs/screenshots/ticker.png" width="230" alt="คำสั่ง bitkub market ticker btc_thb แสดงราคา BTC/THB"> <img src="docs/screenshots/depth.png" width="230" alt="คำสั่ง bitkub market depth แสดง bid และ ask อย่างละ 5 รายการ"> <img src="docs/screenshots/stream.png" width="230" alt="คำสั่ง bitkub stream ticker แสดง ticker แบบ real-time">

**ถ้า go-bitkub ช่วยให้คุณไม่ต้องนั่งอ่านเอกสาร exchange ทั้งบ่าย ช่วย[กด ⭐](https://github.com/dvgamerr-app/go-bitkub/stargazers) ให้หน่อยนะคะ จะได้ช่วยให้นักพัฒนาคริปโตไทยคนอื่นเจอโปรเจกต์นี้**

</div>

## ทำไมต้อง go-bitkub

- **import เดียวใช้ได้ทุก API ของ Bitkub** ทั้ง Market V3, Crypto V4, fiat, user และ WebSocket ใช้ client, error type และ credentials ชุดเดียวกัน
- **ดูข้อมูลตลาดได้โดยไม่ต้องมีบัญชี** คำสั่งและฟังก์ชัน public ไม่ขอ key ไม่ sign และไม่เรียก endpoint เวลาเซิร์ฟเวอร์
- **ยืนยันตัวตนถูกต้องตั้งแต่แรก** คำสั่ง secure จะ sync เวลาเซิร์ฟเวอร์ Bitkub ครั้งเดียว sign ด้วย HMAC-SHA256 และส่ง header ให้ครบ ถ้าไม่ได้ตั้ง key จะแจ้งชัดเจนทันทีแทนที่จะได้ HTTP 400
- **CLI ใช้ใน script ได้ทันที** ทุกคำสั่งพิมพ์ log อ่านง่าย, JSON array หรือ JSON Lines ต่อกับ `jq`, cron job และ dashboard ได้เลย

## ฟีเจอร์

| | ฟีเจอร์ | ได้อะไร |
|---|---|---|
| 📈 | **ข้อมูลตลาด** | symbols, ticker, depth, bids, asks, รายการซื้อขายล่าสุด และแท่งเทียน TradingView |
| 🔴 | **สตรีมสด** | WebSocket ของ trade, ticker และ order book พร้อม reconnect อัตโนมัติ |
| 🧾 | **เทรด** | ตั้ง ยกเลิก และดูคำสั่งซื้อขาย, open orders และประวัติคำสั่ง |
| 💰 | **ยอดคงเหลือ** | wallet และ balance แบบละเอียด รวมยอดที่ถูกจองไว้ |
| 🪙 | **Crypto V4** | ที่อยู่ฝากเหรียญ, ประวัติฝากถอน, เหรียญและเครือข่าย, ค่าชดเชย |
| 🏦 | **Fiat** | บัญชีธนาคาร, ประวัติฝากถอนเงินบาท และการถอนเงินบาท |
| 👤 | **ผู้ใช้** | วงเงินเทรด, trading credits และประวัติ coin convert |
| 🖥️ | **CLI 33 คำสั่ง** | ทุกฟีเจอร์ของ SDK ใช้ผ่าน terminal ได้ 9 คำสั่งไม่ต้องใช้ API key |
| 🧩 | **output 3 รูปแบบ** | `--format text`, `json` หรือ `jsonl` ได้ทุกคำสั่ง |
| 📦 | **binary สำเร็จรูป** | Linux, macOS และ Windows ทั้ง amd64 และ arm64 (Windows มีแค่ amd64) |

<details>
<summary><b>ภาพหน้าจอเพิ่มเติม</b></summary>

<br>

<img src="docs/screenshots/json.png" width="640" alt="output แบบ JSON ของคำสั่ง ticker ที่ส่งต่อให้ jq">

<img src="docs/screenshots/depth.png" width="640" alt="output ของคำสั่ง depth">

<img src="docs/screenshots/stream.png" width="640" alt="output ของสตรีม ticker">

</details>

## เริ่มใช้งาน

### 1. ติดตั้ง CLI

ถ้ามี Go 1.26 ขึ้นไป:

```bash
go install github.com/dvgamerr-app/go-bitkub/cmd/bitkub@latest
```

ถ้าไม่มี Go ให้ดาวน์โหลด binary (ตัวอย่างเป็น Linux x86_64 เครื่องอื่นใช้ `Darwin_arm64`, `Darwin_x86_64` หรือ `Linux_arm64`):

```bash
curl -fsSL https://github.com/dvgamerr-app/go-bitkub/releases/latest/download/go-bitkub_Linux_x86_64.tar.gz | tar -xz bitkub
sudo mv bitkub /usr/local/bin/
```

บน Windows ให้ดาวน์โหลด `go-bitkub_Windows_x86_64.zip` จาก [release ล่าสุด](https://github.com/dvgamerr-app/go-bitkub/releases/latest) แล้ววาง `bitkub.exe` ไว้ใน `PATH`

### 2. ดูข้อมูลตลาด (ไม่ต้องใช้ key)

```bash
bitkub market ticker btc_thb
bitkub market depth btc_thb --limit 5
bitkub stream ticker thb_btc -n 3
```

### 3. เชื่อมบัญชี (ไม่บังคับ)

สร้าง API key ในหน้าจัดการ API ของ Bitkub แล้ว export หรือใส่ในไฟล์ `.env` ในโฟลเดอร์ที่รัน `bitkub`:

```bash
export BTK_APIKEY=REPLACE_ME
export BTK_SECRET=REPLACE_ME

bitkub market balances
```

หรือส่ง `--key` และ `--secret` ในคำสั่งก็ได้ ดูรายละเอียดเพิ่มใน[คู่มือติดตั้ง](INSTALL.md)

## ใช้จาก Go

```bash
go get github.com/dvgamerr-app/go-bitkub
```

```go
tickers, err := market.GetTicker("btc_thb")
if err != nil {
	log.Fatal(err)
}
fmt.Printf("BTC/THB %.2f (%+.2f%%)\n", tickers[0].Last, tickers[0].PercentChange)

bitkub.Initlizer()
balances, err := market.GetBalances()
```

`bitkub.Initlizer()` อ่าน `BTK_APIKEY` และ `BTK_SECRET` จาก environment หรือรับเป็น argument สองตัว ตัวอย่างเต็มอยู่ใน [README ภาษาอังกฤษ](README.md#use-it-from-go) และ[เอกสาร SDK](docs/reference.md)

## สถาปัตยกรรม

```
  your Go app ─┐                    ┌─ https://api.bitkub.com/api/v3/*   market, trading, fiat, user
               ├─▶ go-bitkub SDK ───┼─ https://api.bitkub.com/api/v4/*   crypto deposits and withdrawals
  bitkub CLI ──┘   (bitkub/fetch)   └─ wss://api.bitkub.com/websocket-api   live streams
```

Go 1.26 ใช้ `net/http` มาตรฐาน, `gorilla/websocket` สำหรับสตรีม, Cobra สำหรับ CLI และ zerolog สำหรับ output

## เอกสาร

- [คู่มือติดตั้ง](INSTALL.md)
- [คู่มือ CLI](CLI.md)
- [เอกสาร SDK](docs/reference.md)
- [คู่มือสตรีม](stream/README.md)
- [เอกสาร API ทางการของ Bitkub](https://github.com/bitkub/bitkub-official-api-docs)

## พัฒนาต่อ

```bash
git clone https://github.com/dvgamerr-app/go-bitkub.git
cd go-bitkub
go test ./...
go vet ./...
go build ./...
```

## สนับสนุนโปรเจกต์

- ⭐ [กดดาวให้ repository](https://github.com/dvgamerr-app/go-bitkub/stargazers)
- 🐛 [แจ้งบั๊กหรือขอ endpoint เพิ่ม](https://github.com/dvgamerr-app/go-bitkub/issues)
- 💬 [เข้า Discord](https://discord.gg/QDccF497Mw)

## License

[MIT](LICENSE) go-bitkub เป็นโปรเจกต์ที่ไม่เป็นทางการ ไม่มีส่วนเกี่ยวข้องหรือได้รับการรับรองจาก Bitkub การเทรดคริปโตมีความเสี่ยง ควรทดสอบด้วยจำนวนเงินน้อยๆ ก่อน
