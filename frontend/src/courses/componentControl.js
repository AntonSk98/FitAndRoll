export class ComponentControl {
    constructor() {
        this.showCourseOverview = true;
        this.defineNewCourseComponent = false;
        this.updateCourseComponent = false;
        this.courseParticipantsComponent = false;
        this.courseParticipantsHistoryComponent = false;
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
    }

    showcourseParticipantsHistoryComponent() {
        this.courseParticipantsHistoryComponent = true;
        this.showCourseOverview = false;
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