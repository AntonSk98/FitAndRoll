export namespace archive {
	
	export class ArchivedEntryDto {
	    id: number;
	    name: string;
	    archivedAt: string;
	
	    static createFrom(source: any = {}) {
	        return new ArchivedEntryDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.archivedAt = source["archivedAt"];
	    }
	}
	export class FindArchivedEntryParams {
	    name?: string;
	
	    static createFrom(source: any = {}) {
	        return new FindArchivedEntryParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	    }
	}

}

export namespace common {
	
	export class PageParams {
	    currentPage: number;
	    pageSize: number;
	
	    static createFrom(source: any = {}) {
	        return new PageParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.currentPage = source["currentPage"];
	        this.pageSize = source["pageSize"];
	    }
	}
	export class Page[fit_and_roll/backend/archive.ArchivedEntryDto] {
	    data: archive.ArchivedEntryDto[];
	    total: number;
	    page: number;
	    size: number;
	
	    static createFrom(source: any = {}) {
	        return new Page[fit_and_roll/backend/archive.ArchivedEntryDto](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = this.convertValues(source["data"], archive.ArchivedEntryDto);
	        this.total = source["total"];
	        this.page = source["page"];
	        this.size = source["size"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Page[fit_and_roll/backend/courseattendance.CourseAttendanceDto] {
	    data: courseattendance.CourseAttendanceDto[];
	    total: number;
	    page: number;
	    size: number;
	
	    static createFrom(source: any = {}) {
	        return new Page[fit_and_roll/backend/courseattendance.CourseAttendanceDto](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = this.convertValues(source["data"], courseattendance.CourseAttendanceDto);
	        this.total = source["total"];
	        this.page = source["page"];
	        this.size = source["size"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Page[fit_and_roll/backend/courses.CourseDto] {
	    data: courses.CourseDto[];
	    total: number;
	    page: number;
	    size: number;
	
	    static createFrom(source: any = {}) {
	        return new Page[fit_and_roll/backend/courses.CourseDto](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = this.convertValues(source["data"], courses.CourseDto);
	        this.total = source["total"];
	        this.page = source["page"];
	        this.size = source["size"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Page[fit_and_roll/backend/participants.ParticipantDto] {
	    data: participants.ParticipantDto[];
	    total: number;
	    page: number;
	    size: number;
	
	    static createFrom(source: any = {}) {
	        return new Page[fit_and_roll/backend/participants.ParticipantDto](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = this.convertValues(source["data"], participants.ParticipantDto);
	        this.total = source["total"];
	        this.page = source["page"];
	        this.size = source["size"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Page[fit_and_roll/backend/statistics.StatisticsDto] {
	    data: statistics.StatisticsDto[];
	    total: number;
	    page: number;
	    size: number;
	
	    static createFrom(source: any = {}) {
	        return new Page[fit_and_roll/backend/statistics.StatisticsDto](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = this.convertValues(source["data"], statistics.StatisticsDto);
	        this.total = source["total"];
	        this.page = source["page"];
	        this.size = source["size"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class TimeRangeFilter {
	    // Go type: time
	    from?: any;
	    // Go type: time
	    to?: any;
	
	    static createFrom(source: any = {}) {
	        return new TimeRangeFilter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.from = this.convertValues(source["from"], null);
	        this.to = this.convertValues(source["to"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace courseattendance {
	
	export class CourseAttendanceDto {
	    id: number;
	    fullname: string;
	    course: string;
	    attendedAt: string;
	    attendanceType: string;
	
	    static createFrom(source: any = {}) {
	        return new CourseAttendanceDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.fullname = source["fullname"];
	        this.course = source["course"];
	        this.attendedAt = source["attendedAt"];
	        this.attendanceType = source["attendanceType"];
	    }
	}
	export class CourseAttendanceParameters {
	    courseId?: number;
	    fullname?: string;
	    course?: string;
	    excludeArchivedCourse: boolean;
	    excludeTrialAttendance: boolean;
	    excludeNoMemberCard: boolean;
	    excludeWithMemberCard: boolean;
	    attendedRange?: common.TimeRangeFilter;
	
	    static createFrom(source: any = {}) {
	        return new CourseAttendanceParameters(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.courseId = source["courseId"];
	        this.fullname = source["fullname"];
	        this.course = source["course"];
	        this.excludeArchivedCourse = source["excludeArchivedCourse"];
	        this.excludeTrialAttendance = source["excludeTrialAttendance"];
	        this.excludeNoMemberCard = source["excludeNoMemberCard"];
	        this.excludeWithMemberCard = source["excludeWithMemberCard"];
	        this.attendedRange = this.convertValues(source["attendedRange"], common.TimeRangeFilter);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace courses {
	
	export class ScheduleDto {
	    day: string;
	    time: string;
	
	    static createFrom(source: any = {}) {
	        return new ScheduleDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.day = source["day"];
	        this.time = source["time"];
	    }
	}
	export class CourseDetailsDto {
	    id: number;
	    name: string;
	    description: string;
	    schedules: ScheduleDto[];
	
	    static createFrom(source: any = {}) {
	        return new CourseDetailsDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.schedules = this.convertValues(source["schedules"], ScheduleDto);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class CourseDto {
	    id: number;
	    name: string;
	    description: string;
	    schedules: ScheduleDto[];
	
	    static createFrom(source: any = {}) {
	        return new CourseDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.schedules = this.convertValues(source["schedules"], ScheduleDto);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ScheduleEntryRequest {
	    day: string;
	    time: string;
	
	    static createFrom(source: any = {}) {
	        return new ScheduleEntryRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.day = source["day"];
	        this.time = source["time"];
	    }
	}
	export class CourseRequest {
	    name: string;
	    description: string;
	    schedules: ScheduleEntryRequest[];
	
	    static createFrom(source: any = {}) {
	        return new CourseRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	        this.schedules = this.convertValues(source["schedules"], ScheduleEntryRequest);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class FindCourseParams {
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new FindCourseParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	    }
	}
	
	
	export class UpdateCourseRequest {
	    name: string;
	    description: string;
	    schedules: ScheduleEntryRequest[];
	    id?: number;
	
	    static createFrom(source: any = {}) {
	        return new UpdateCourseRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	        this.schedules = this.convertValues(source["schedules"], ScheduleEntryRequest);
	        this.id = source["id"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace membercardattendance {
	
	export class CourseAttendanceCommand {
	    memberCard?: number;
	    course: number;
	    participant: number;
	    attendanceType: string;
	
	    static createFrom(source: any = {}) {
	        return new CourseAttendanceCommand(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.memberCard = source["memberCard"];
	        this.course = source["course"];
	        this.participant = source["participant"];
	        this.attendanceType = source["attendanceType"];
	    }
	}
	export class MemberCardHistoryEntry {
	    course: string;
	    attendedAt: string;
	
	    static createFrom(source: any = {}) {
	        return new MemberCardHistoryEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.course = source["course"];
	        this.attendedAt = source["attendedAt"];
	    }
	}

}

export namespace participants {
	
	export class FindParticipantsParams {
	    name: string;
	    surname: string;
	    group: string;
	    withActiveCard: boolean;
	
	    static createFrom(source: any = {}) {
	        return new FindParticipantsParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.surname = source["surname"];
	        this.group = source["group"];
	        this.withActiveCard = source["withActiveCard"];
	    }
	}
	export class MemberCardInfo {
	    id: number;
	    capacity: number;
	    issuedAt: string;
	    expiredAt?: string;
	
	    static createFrom(source: any = {}) {
	        return new MemberCardInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.capacity = source["capacity"];
	        this.issuedAt = source["issuedAt"];
	        this.expiredAt = source["expiredAt"];
	    }
	}
	export class ParticipantCommand {
	    id?: number;
	    name: string;
	    surname: string;
	    group: string;
	
	    static createFrom(source: any = {}) {
	        return new ParticipantCommand(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.surname = source["surname"];
	        this.group = source["group"];
	    }
	}
	export class ParticipantDto {
	    id?: number;
	    name: string;
	    surname: string;
	    group: string;
	
	    static createFrom(source: any = {}) {
	        return new ParticipantDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.surname = source["surname"];
	        this.group = source["group"];
	    }
	}

}

export namespace statistics {
	
	export class StatisticsDto {
	    name: string;
	    withMemberCard: number;
	    trialAttendance: number;
	    withoutMemberCard: number;
	
	    static createFrom(source: any = {}) {
	        return new StatisticsDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.withMemberCard = source["withMemberCard"];
	        this.trialAttendance = source["trialAttendance"];
	        this.withoutMemberCard = source["withoutMemberCard"];
	    }
	}
	export class StatisticsParams {
	    name?: string;
	    attendedRange?: common.TimeRangeFilter;
	
	    static createFrom(source: any = {}) {
	        return new StatisticsParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.attendedRange = this.convertValues(source["attendedRange"], common.TimeRangeFilter);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

