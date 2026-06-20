package main

import (
	"fmt"
	"os"
	"strings"
)

func printUsage() {
	fmt.Println("社区煤气站气瓶管理工具")
	fmt.Println()
	fmt.Println("用法:")
	fmt.Println("  gasstation add-cylinder <编号> --type <5kg/15kg/50kg> --owner <自有/客户自带>  登记气瓶")
	fmt.Println("  gasstation fill <编号> --date <日期> --operator <操作人>                      充装记录")
	fmt.Println("  gasstation lend <编号> --customer <客户> --phone <电话> --date <日期>          借出气瓶")
	fmt.Println("  gasstation return <编号> --date <日期>                                        归还气瓶")
	fmt.Println("  gasstation monthly --month <YYYY-MM>                                          月度汇总")
	fmt.Println("  gasstation status                                                              状态汇总")
}

func parseFlags(args []string) (map[string]string, []string) {
	flags := map[string]string{}
	var positional []string
	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "--") {
			key := args[i][2:]
			if idx := strings.Index(key, "="); idx >= 0 {
				flags[key[:idx]] = key[idx+1:]
			} else if i+1 < len(args) && !strings.HasPrefix(args[i+1], "--") {
				flags[key] = args[i+1]
				i++
			} else {
				flags[key] = ""
			}
		} else {
			positional = append(positional, args[i])
		}
	}
	return flags, positional
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	data, err := loadData()
	if err != nil {
		fmt.Fprintf(os.Stderr, "加载数据失败: %v\n", err)
		os.Exit(1)
	}

	needSave := true

	switch os.Args[1] {
	case "add-cylinder":
		flags, positional := parseFlags(os.Args[2:])
		if len(positional) < 1 {
			fmt.Fprintln(os.Stderr, "错误: 请提供气瓶编号")
			os.Exit(1)
		}
		err = HandleAddCylinder(data, positional[0], flags["type"], flags["owner"])

	case "fill":
		flags, positional := parseFlags(os.Args[2:])
		if len(positional) < 1 {
			fmt.Fprintln(os.Stderr, "错误: 请提供气瓶编号")
			os.Exit(1)
		}
		err = HandleFill(data, positional[0], flags["date"], flags["operator"])

	case "lend":
		flags, positional := parseFlags(os.Args[2:])
		if len(positional) < 1 {
			fmt.Fprintln(os.Stderr, "错误: 请提供气瓶编号")
			os.Exit(1)
		}
		err = HandleLend(data, positional[0], flags["customer"], flags["phone"], flags["date"])

	case "return":
		flags, positional := parseFlags(os.Args[2:])
		if len(positional) < 1 {
			fmt.Fprintln(os.Stderr, "错误: 请提供气瓶编号")
			os.Exit(1)
		}
		err = HandleReturn(data, positional[0], flags["date"])

	case "monthly":
		flags, _ := parseFlags(os.Args[2:])
		HandleMonthly(data, flags["month"])
		needSave = false

	case "status":
		HandleStatus(data)
		needSave = false

	default:
		fmt.Fprintf(os.Stderr, "未知命令: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}

	if needSave {
		if err = saveData(data); err != nil {
			fmt.Fprintf(os.Stderr, "保存数据失败: %v\n", err)
			os.Exit(1)
		}
	}
}
