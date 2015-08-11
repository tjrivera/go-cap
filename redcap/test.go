package redcap

var project = NewRedcapProject("https://redcap.vanderbilt.edu/api/", "8E66DB6844D58E990075AFB51658A002", true)

// Uninitialized Project
var project_uninit = NewRedcapProject("https://redcap.vanderbilt.edu/api/", "8E66DB6844D58E990075AFB51658A002", false)
