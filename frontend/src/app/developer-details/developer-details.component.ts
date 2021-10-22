import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Developer } from '../developer';
import { DeveloperService } from '../developer.service';
import { Location } from '@angular/common';

@Component({
  selector: 'app-developer-details',
  templateUrl: './developer-details.component.html',
  styleUrls: ['./developer-details.component.css'],
})
export class DeveloperDetailsComponent implements OnInit {
  developer: Developer | undefined;

  constructor(
    private route: ActivatedRoute,
    private developerService: DeveloperService,
    private location: Location
  ) {}

  ngOnInit(): void {
    this.getDeveloper();
  }

  getDeveloper(): void {
    const id = Number(this.route.snapshot.paramMap.get('id'));
    this.developerService
      .getDeveloper(id)
      .subscribe((developer) => (this.developer = developer));
  }

  goBack(): void {
    this.location.back();
  }

  save(): void {
    if (this.developer) {
      this.developerService.updateDeveloper(this.developer).subscribe();
    }
  }
}
