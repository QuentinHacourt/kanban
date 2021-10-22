import { Component, OnInit } from '@angular/core';
import { Developer } from '../developer';
import { DeveloperService } from '../developer.service';

@Component({
  selector: 'app-developers',
  templateUrl: './developers.component.html',
  styleUrls: ['./developers.component.css'],
})
export class DevelopersComponent implements OnInit {
  developers: Developer[] = [];

  constructor(private developerService: DeveloperService) {}

  ngOnInit(): void {
    this.getDevelopers();
  }

  getDevelopers(): void {
    this.developerService
      .getDevelopers()
      .subscribe((developers) => (this.developers = developers));
  }

  delete(developer: Developer): void {
    this.developers = this.developers.filter((h) => h !== developer);
    this.developerService.deleteDeveloper(developer.id).subscribe();
  }
}
