// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {membercardattendance} from '../models';
import {participants} from '../models';

export function AttendCourse(arg1:membercardattendance.CourseAttendanceCommand):Promise<void>;

export function FindActiveMemberCards(arg1:number):Promise<Array<participants.MemberCardInfo>>;

export function LoadMemberCardCourseHistory(arg1:number,arg2:number):Promise<Array<membercardattendance.MemberCardHistoryEntry>>;
