package main

import "fmt"
import "sort"

type Task struct {
    ID   string
    ServiceID string
    NodeID string
}



func main(){
    var task2, task3 Task;

    task2.ID = "taskID2"
    task2.ServiceID = "ServiceID2"
    task2.NodeID = "NodeID2"

    task3.ID = "taskID3"
    task3.ServiceID = "ServiceID3"
    task3.NodeID = "NodeID3"

    task1 := Task{
        ID: "taskID1",
        ServiceID: "ServiceID1",
        NodeID: "NodeID1",
    }

    var taskGroup map[string]Task;
    taskGroup = make(map[string]Task)

    taskGroup["taskID1"] = task1
    taskGroup["taskID3"] = task3
    taskGroup["taskID2"] = task2

    // 用于记录serviceID，用于排序
    var serviceID = []string{}

    // 用于标记task是否指派
    var serviceID_flag map[string]string;
    serviceID_flag = make(map[string]string)

    //遍历task，收集serviceID
    for taskID := range taskGroup{
        // 收集serviceID
        serviceID = append(serviceID, taskGroup[taskID].ServiceID)
        // 默认未指派
        serviceID_flag[taskID] = "false"
        // 测试输出结果
        fmt.Println(taskID, taskGroup[taskID])
    }

    // 测试输出结果
    fmt.Println("任务指派前，serviceID_flag的值")
    for taskID := range serviceID_flag{
        fmt.Println(taskID, serviceID_flag[taskID])
    }

    fmt.Println("排序前：", serviceID)
    // 排序
    sort.Sort(sort.StringSlice(serviceID))
    fmt.Println("排序后：", serviceID)


    for _, value := range serviceID{
        assign_task_ID := ""
        for taskID := range taskGroup{
            if value == taskGroup[taskID].ServiceID && serviceID_flag[taskID] == "false"{
                assign_task_ID = taskID
                serviceID_flag[taskID] = "true"
                break
            }
        }
        fmt.Println(assign_task_ID, taskGroup[assign_task_ID])
    }

    fmt.Println("任务指派后，serviceID_flag的值")
    for taskID := range serviceID_flag{
        fmt.Println(taskID, serviceID_flag[taskID])
    }
}