package analysis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

// 顺序信息结构
type ExecutionOrder struct {
	Function string `json:"function"` // 格式：<ContractName>::<FunctionName>
	Order    int    `json:"order"`
}

// 全局顺序映射表
var executionOrderMap map[string]int

// 顺序信息初始化（只加载一次）
func init() {
	executionOrderMap = make(map[string]int)

	// 假设路径固定，sorter 的输出是一个 JSON 格式的列表
	data, err := ioutil.ReadFile("/Users/lvsirui/fabric-tape/SEFabric/fabric/core/chaincode/analysis/output_order.json")
	if err != nil {
		fmt.Printf("Error reading order file: %v\n", err)
		return
	}

	var orders []ExecutionOrder
	err = json.Unmarshal(data, &orders)
	if err != nil {
		fmt.Printf("Error parsing order JSON: %v\n", err)
		return
	}

	for _, entry := range orders {
		executionOrderMap[entry.Function] = entry.Order
	}
}

// Invoke 包裹函数，在这里设置事件记录执行顺序
func Invoke(stub shim.ChaincodeStubInterface, function string, args []string) pb.Response {
	// 拼接 key，例如 smallbank::Amalgamate
	fullFunction := formatFunctionKey(function)

	// 查找顺序号
	order, ok := executionOrderMap[fullFunction]
	if !ok {
		return shim.Error(fmt.Sprintf("No execution order found for function %s", fullFunction))
	}

	// 设置事件，所有节点都能获取顺序信息
	orderStr := fmt.Sprintf("%d", order)
	err := stub.SetEvent("ExecutionOrder", []byte(orderStr))
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to set execution order event: %v", err))
	}

	// TODO: 实际调用被测试合约函数
	// 你需要根据你的架构调用对应的 chaincode handler
	// 这里只是一个占位：
	fmt.Printf("Executing function %s with order %d\n", fullFunction, order)

	// 返回成功响应
	return shim.Success([]byte(fmt.Sprintf("Executed %s with order %d", fullFunction, order)))
}

// 将 function 参数标准化成 key，例如 Amalgamate => smallbank::Amalgamate
func formatFunctionKey(function string) string {
	// 你可以根据传参情况自定义这个逻辑
	// 例如直接传入完整的 "smallbank::Amalgamate"
	if strings.Contains(function, "::") {
		return function
	}
	// 默认合约名为 smallbank
	return "smallbank::" + function
}
