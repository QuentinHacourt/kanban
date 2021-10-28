import { Component, OnInit } from '@angular/core';
import { ProjectInput } from '../project';
import { ProjectService } from '../project.service';
import { Team } from '../team';
import { TeamService } from '../team.service';

@Component({
  selector: 'app-create-project',
  templateUrl: './create-project.component.html',
  styleUrls: ['./create-project.component.css'],
})
export class CreateProjectComponent implements OnInit {
  teams: Team[] = [];

  constructor(
    private projectService: ProjectService,
    private teamService: TeamService
  ) {}

  ngOnInit(): void {
    this.getTeams();
  }

  addProject(title: string, description: string, teamName: string): void {
    if (!title || !description) {
      return;
    }
    const projectInput: ProjectInput = {
      title: title,
      description: description,
      team_name: teamName,
    };
    this.projectService.addProject(projectInput).subscribe();
  }

  getTeams(): void {
    this.teamService.getTeams().subscribe((teams) => (this.teams = teams));
  }
}
