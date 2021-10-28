import { Component, OnInit } from '@angular/core';
import { DeveloperInput } from '../developer';
import { DeveloperService } from '../developer.service';
import { Team } from '../team';
import { TeamService } from '../team.service';

@Component({
  selector: 'app-register-developer',
  templateUrl: './register-developer.component.html',
  styleUrls: ['./register-developer.component.css'],
})
export class RegisterDeveloperComponent implements OnInit {
  teams: Team[] = [];

  constructor(
    private developerService: DeveloperService,
    private teamService: TeamService
  ) {}

  ngOnInit(): void {
    this.getTeams();
  }

  registerDeveloper(
    user_name: string,
    password: string,
    teamName: string
  ): void {
    if (!user_name || !password || !teamName) {
      return;
    }
    const developerInput: DeveloperInput = {
      user_name: user_name,
      password: password,
      team_name: teamName,
    };
    this.developerService.addDeveloper(developerInput).subscribe();
  }

  getTeams(): void {
    this.teamService.getTeams().subscribe((teams) => (this.teams = teams));
  }
}
