# sap-api-integrations-service-order-reads
sap-api-integrations-service-order-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で サービス指図データを取得するマイクロサービスです。    
sap-api-integrations-service-order-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-service-order-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_SERVICE_ORDER_SRV_0001/overview  

## 動作環境  
sap-api-integrations-service-order-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）　　

## クラウド環境での利用
sap-api-integrations-service-order-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-service-order-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_SERVICE_ORDER_SRV_0001/overview  
* APIサービス名(=baseURL): API_SERVICE_ORDER_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-service-order-reads には、次の API をコールするためのリソースが含まれています。  

* A_SalesOrder（サービス指図 - ヘッダ）※サービス指図の詳細データを取得するために、ToHeaderConfirmation、ToHeaderDefect、ToHeaderPersonResponsible、ToHeaderReferenceObject、ToItem、ToItemPricingElementと合わせて利用されます。
* A_ServiceOrderItem（サービス指図 - 明細）※受注明細の詳細データを取得するために、ToItemPricingElement、と合わせて利用されます。
* ToHeaderConfirmation（サービス指図 - 確認）
* ToHeaderDefect（サービス指図 - 不良）
* ToHeaderPersonResponsible（サービス指図 - 責任者）
* ToHeaderReferenceObject（サービス指図 - 参照対象）
* ToItem（サービス指図 - 明細）
* ToItemPricingElement（サービス指図 - 明細価格条件）
* A_ServiceOrderItem（サービス指図 - 明細）※受注明細の詳細データを取得するために、ToItemPricingElement、と合わせて利用されます。

## API への 値入力条件 の 初期値
sap-api-integrations-service-order-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.ServiceOrder.ServiceOrder（サービス指図番号）
* inoutSDC.ServiceOrder.ServiceOrderItem.ServiceOrderItem（サービス指図明細）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header", "Item" が指定されています。

```
	"api_schema": "sap.s4.beh.serviceorder.v1.ServiceOrder.Created.v1",
	"accepter": ["Header"],
	"service_order": "4000000000",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "sap.s4.beh.serviceorder.v1.ServiceOrder.Created.v1",
	"accepter": ["All"],
	"service_order": "4000000000",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetServiceOrder(serviceOrder, serviceOrderItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(serviceOrder)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(serviceOrder, serviceOrderItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```
## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP サービス指図 の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ServiceOrder" ～ "to_Item" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-service-order-reads/SAP_API_Caller/caller.go#L58",
	"function": "sap-api-integrations-service-order-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"ServiceOrder": "4000000000",
			"ServiceOrderType": "RPO1",
			"ServiceOrderUUID": "fa163e6c-c3e1-1edb-a491-b1f51f15989f",
			"ServiceOrderDescription": "Printer Professional IHR1911",
			"ServiceObjectType": "BUS2000116",
			"Language": "EN",
			"ServiceDocumentPriority": "0",
			"RequestedServiceStartDateTime": "/Date(1616976000000+0000)/",
			"RequestedServiceEndDateTime": "/Date(1617235200000+0000)/",
			"ServiceDocChangedDateTime": "/Date(1617022513000+0000)/",
			"PurchaseOrderByCustomer": "",
			"CustomerPurchaseOrderDate": "",
			"ServiceOrderIsReleased": "",
			"ServiceOrderIsCompleted": "X",
			"ServiceOrderIsRejected": "",
			"SalesOrganization": "1710",
			"DistributionChannel": "10",
			"Division": "00",
			"SalesOffice": "",
			"SalesGroup": "",
			"SoldToParty": "17100001",
			"ShipToParty": "17100001",
			"BillToParty": "17100001",
			"PayerParty": "17100001",
			"ContactPerson": "",
			"ServiceDocGrossAmount": "36.40",
			"ServiceDocNetAmount": "35.00",
			"ServiceDocTaxAmount": "1.40",
			"TransactionCurrency": "USD",
			"ReferenceServiceRequest": "",
			"ReferenceServiceContract": "",
			"ReferenceServiceOrder": "",
			"SrvcOrdCreditStatus": "",
			"ServiceOrderRejectionReason": "",
			"to_Confirmation": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_SERVICE_ORDER_SRV/A_ServiceOrder('4000000000')/to_Confirmation",
			"to_Defect": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_SERVICE_ORDER_SRV/A_ServiceOrder('4000000000')/to_Defect",
			"to_PersonResponsible": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_SERVICE_ORDER_SRV/A_ServiceOrder('4000000000')/to_PersonResponsible",
			"to_ReferenceObject": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_SERVICE_ORDER_SRV/A_ServiceOrder('4000000000')/to_ReferenceObject",
			"to_Item": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_SERVICE_ORDER_SRV/A_ServiceOrder('4000000000')/to_Item"
		}
	],
	"time": "2021-12-19T14:16:34.713317+09:00"
}

```
