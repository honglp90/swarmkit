package main

import "fmt"
import "sort"
import "unsafe"

type Resources struct {
    MemoryBytes int64
}

type ResourceRequirements struct {
    Limits       *Resources
    Reservations *Resources
}

type TaskSpec struct {
    Resources *ResourceRequirements
    // ResourceReferences []ResourceReference
}

type Task struct {
    ID   string
    ServiceID string
    NodeID string
    Spec TaskSpec
}


func main(){
    task1 := Task{
        ID: "taskID1",
        ServiceID: "ServiceID1",
        NodeID: "NodeID1",
        Spec: TaskSpec{
            Resources: &ResourceRequirements{
                Limits: &Resources{
                    MemoryBytes: 1024000000,
                },
                Reservations: &Resources{
                    MemoryBytes: 1024000000,
                },
            },
        },
    }

    task2 := Task{
        ID: "taskID2",
        ServiceID: "ServiceID2",
        NodeID: "NodeID2",
        Spec: TaskSpec{
            Resources: &ResourceRequirements{
                Limits: &Resources{
                    MemoryBytes: 4096000000,
                },
                Reservations: &Resources{
                    MemoryBytes: 1024000000,
                },
            },
        },
    }

    task3 := Task{
        ID: "taskID3",
        ServiceID: "ServiceID3",
        NodeID: "NodeID3",
        Spec: TaskSpec{
            Resources: &ResourceRequirements{
                Limits: &Resources{
                    MemoryBytes: 2048000000,
                },
                Reservations: &Resources{
                    MemoryBytes: 1024000000,
                },
            },
        },
    }

    task4 := Task{
        ID: "taskID4",
        ServiceID: "ServiceID4",
        NodeID: "NodeID4",
        Spec: TaskSpec{
            Resources: &ResourceRequirements{
                Limits: &Resources{
                    MemoryBytes: 1024000000,
                },
                Reservations: &Resources{
                    MemoryBytes: 1024000000,
                },
            },
        },
    }

    var taskGroup map[string]*Task;
    taskGroup = make(map[string]*Task)

    taskGroup["taskID1"] = &task1
    taskGroup["taskID2"] = &task2
    taskGroup["taskID3"] = &task3
    taskGroup["taskID4"] = &task4

    // 用于记录内存，用于排序
    var memory_list = []int{}
    var memory int64

    // 用于标记task是否指派
    var task_assign_flag map[string]string;
    task_assign_flag = make(map[string]string)

    //遍历task，收集memory_list
    for taskID := range taskGroup{
        // 收集memory_list
        memory = (*(*(*taskGroup[taskID]).Spec.Resources).Limits).MemoryBytes
        memory_list = append(memory_list, *(*int)(unsafe.Pointer(&memory)))
        // 默认未指派
        task_assign_flag[taskID] = "false"
    }

    // 测试输出结果
    fmt.Println("任务指派前，task_assign_flag的值")
    for taskID := range task_assign_flag{
        fmt.Println(taskID, task_assign_flag[taskID])
    }

    fmt.Println("排序前：", memory_list)
    // 降序排列
    sort.Sort(sort.Reverse(sort.IntSlice(memory_list)))
    fmt.Println("排序后：", memory_list)


    for _, value := range memory_list{
        taskID := ""
        var t *Task
        for assign_task_ID, assign_task := range taskGroup{
            memory = (*(*(*taskGroup[assign_task_ID]).Spec.Resources).Limits).MemoryBytes
            if value == *(*int)(unsafe.Pointer(&memory)) && task_assign_flag[assign_task_ID] == "false"{
                taskID = assign_task_ID
                t = assign_task
                task_assign_flag[assign_task_ID] = "true"
                break
            }
        }
        fmt.Printf("%+v\n", *taskGroup[taskID])
        fmt.Printf("%s, %+v, %+v\n", taskID, (*(*(*taskGroup[taskID]).Spec.Resources).Limits).MemoryBytes, (*(*t.Spec.Resources).Limits).MemoryBytes)
    }

    fmt.Println("任务指派后，task_assign_flag的值")
    for taskID := range task_assign_flag{
        fmt.Println(taskID, task_assign_flag[taskID])
    }
}