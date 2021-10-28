import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Project } from '../project';
import { ProjectService } from '../project.service';
import { Location } from '@angular/common';
import { Team } from '../team';
import { TeamService } from '../team.service';

@Component({
  selector: 'app-project-details',
  templateUrl: './project-details.component.html',
  styleUrls: ['./project-details.component.css'],
})
export class ProjectDetailsComponent implements OnInit {
  project: Project | undefined;
  teams: Team[] = [];

  constructor(
    private route: ActivatedRoute,
    private projectService: ProjectService,
    private teamService: TeamService,
    private location: Location
  ) {}

  ngOnInit(): void {
    this.getProject();
    this.getTeams();
  }

  getProject(): void {
    const id = Number(this.route.snapshot.paramMap.get('id'));
    this.projectService
      .getProject(id)
      .subscribe((project) => (this.project = project));
  }

  goBack(): void {
    this.location.back();
  }

  save(): void {
    if (this.project) {
      this.projectService.updateProject(this.project).subscribe();
    }
  }

  getTeams(): void {
    this.teamService.getTeams().subscribe((teams) => (this.teams = teams));
  }
}
