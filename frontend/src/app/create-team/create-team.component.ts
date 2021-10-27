import { Component, OnInit } from '@angular/core';
import { TeamInput } from '../team';
import { TeamService } from '../team.service';

@Component({
  selector: 'app-create-team',
  templateUrl: './create-team.component.html',
  styleUrls: ['./create-team.component.css'],
})
export class CreateTeamComponent implements OnInit {
  constructor(private teamService: TeamService) {}

  ngOnInit(): void {}

  addTeam(name: string): void {
    if (!name) {
      return;
    }
    const teamInput: TeamInput = {
      name: name,
    };
    this.teamService.addTeam(teamInput).subscribe();
  }
}
