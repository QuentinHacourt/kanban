import { Component, OnInit } from '@angular/core';
import { Project } from '../project';
import { ProjectService } from '../project.service';

@Component({
  selector: 'app-projects',
  templateUrl: './projects.component.html',
  styleUrls: ['./projects.component.css'],
})
export class ProjectsComponent implements OnInit {
  projects: Project[] = [];

  constructor(private projectService: ProjectService) {}

  ngOnInit(): void {
    this.getProjects();
  }

  getProjects(): void {
    this.projectService
      .getProjects()
      .subscribe((projects) => (this.projects = projects));
  }

  delete(developer: Project): void {
    this.projects = this.projects.filter((h) => h !== developer);
    this.projectService.deleteProject(developer.id).subscribe();
  }
}
