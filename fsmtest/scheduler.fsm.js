// fsm-config: {"rankdir": "LR"}

// fsm-begin  <------!important

let SchedulerFsm = {
    name: "Scheduler",
    initial: 'Idle',

    events: [
        {name: 'AddTask', from: 'Idle', to: 'Assigned'},
        {name: 'StartCmd', from: 'Assigned', to: 'Processing'},
        {name: 'StartCmd', from: 'AssignedMaintenance', to: 'ProcessingMaintenance'},
        {name: 'DoneCmd', from: 'Processing', to: 'Completed'},
        {name: 'DoneCmd', from: 'ProcessingMaintenance', to: 'CompletedMaintenance'},
        {name: 'AbortCmd', from: 'Assigned', to: 'Abort'},
        {name: 'AbortCmd', from: 'Processing', to: 'Abort'},
        {name: 'AbortCmd', from: 'AssignedMaintenance', to: 'AbortMaintenance'},
        {name: 'AbortCmd', from: 'ProcessingMaintenance', to: 'AbortMaintenance'},
        {name: 'Operation', from: 'Assigned', to: 'AssignedMaintenance'},
        {name: 'Operation', from: 'Processing', to: 'ProcessingMaintenance'},
        {name: 'Maintenance', from: 'CompletedMaintenance', to: 'Maintenance'},
        {name: 'Maintenance', from: 'AbortMaintenance', to: 'Maintenance'},
        {name: 'ClearTask', from: 'Completed', to: 'Idle'},
        {name: 'ClearTask', from: 'Abort', to: 'Idle'},
        {name: 'Repair', from: 'Maintenance', to: 'Idle'},
    ],
    states: [
        {name: 'Idle', color: 'green', comments:"initial state"}
    ]
}

// fsm-end    <------!important