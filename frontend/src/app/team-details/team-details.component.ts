import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Team } from '../team';
import { TeamService } from '../team.service';
import { Location } from '@angular/common';

@Component({
  selector: 'app-team-details',
  templateUrl: './team-details.component.html',
  styleUrls: ['./team-details.component.css'],
})
export class TeamDetailsComponent implements OnInit {
  team: Team | undefined;

  constructor(
    private route: ActivatedRoute,
    private teamService: TeamService,
    private location: Location
  ) {}

  ngOnInit(): void {
    this.getTeam();
  }

  getTeam(): void {
    const id = Number(this.route.snapshot.paramMap.get('id'));
    this.teamService.getTeam(id).subscribe((team) => (this.team = team));
  }

  goBack(): void {
    this.location.back();
  }

  save(): void {
    if (this.team) {
      this.teamService.updateTeam(this.team).subscribe();
    }
  }
}
