export class ComponentControl {
    constructor() {
        this.showCourseOverview = true;
        this.defineNewCourseComponent = false;
        this.updateCourseComponent = false;
        this.courseParticipantsComponent = false;
        this.courseParticipantsHistoryComponent = false;
        this.allAttendanceHistoryComponent = false;
    }

    showDefineNewCourseComponent() {
        this.defineNewCourseComponent = true;
        this.showCourseOverview = false;
        return this;
    }

    showUpdateCourseComponent() {
        this.updateCourseComponent = true;
        this.showCourseOverview = false;
        return this;
    }

    showCourseParticipantComponent() {
        this.courseParticipantsComponent = true;
        this.showCourseOverview = false;
        return this;
    }

    showCourseParticipantsHistoryComponent() {
        this.courseParticipantsHistoryComponent = true;
        this.showCourseOverview = false;
        return this;
    }

    showAllAttendanceHistoryComponent() {
        this.allAttendanceHistoryComponent = true;
        this.showCourseOverview = false;
        return this;
    }

    resetComponentControl() {
        this.defineNewCourseComponent = false;
        this.updateCourseComponent = false;
        this.courseParticipantsComponent = false;
        this.courseParticipantsHistoryComponent = false;
        this.showCourseOverview = true;
        return this;
    }

}