/*
 * === This file is part of ALICE O² ===
 *
 * Copyright 2020 CERN and copyright holders of ALICE O².
 * Author: Miltiadis Alexis <miltiadis.alexis@cern.ch>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * In applying this license CERN does not waive the privileges and
 * immunities granted to it by virtue of its status as an
 * Intergovernmental Organization or submit itself to any jurisdiction.
 */

package event

type TaskMessage interface {
	GetName() string
	GetTaskId() string
	GetTaskPID() int
}

type TaskMessageBase struct {
	Name        string `json:"name"`
	TaskId      string `json:"taskId"`
	PID         int32  `json:"pid"`
	MessageType string `json:"_messageType"`
}

func (tmb *TaskMessageBase) GetName() string {
	return tmb.Name
}

func (tmb *TaskMessageBase) GetTaskId() string {
	return tmb.TaskId
}

func (tmb *TaskMessageBase) GetTaskPID() int {
	return int(tmb.PID)
}

func NewTaskMessage(name, id string, pid int32) (tm TaskMessage) {
	tm = &TaskMessageBase{
		Name:        name,
		TaskId:      id,
		PID:         pid,
		MessageType: "TaskMessage",
	}
	return tm
}
