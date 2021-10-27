import { Component, OnInit } from '@angular/core';
import { Story } from '../story';
import { StoryService } from '../story.service';
import { Location } from '@angular/common';
import { ActivatedRoute } from '@angular/router';
import { DeveloperService } from '../developer.service';
import { Developer } from '../developer';
import { Project } from '../project';
import { ProjectService } from '../project.service';

@Component({
  selector: 'app-story-detail',
  templateUrl: './story-detail.component.html',
  styleUrls: ['./story-detail.component.css'],
})
export class StoryDetailComponent implements OnInit {
  story: Story | undefined;
  developers: Developer[] = [];
  projects: Project[] = [];

  constructor(
    private route: ActivatedRoute,
    private storyService: StoryService,
    private developerService: DeveloperService,
    private projectService: ProjectService,
    private location: Location
  ) {}

  ngOnInit(): void {
    this.getStory();
    this.getDevelopers();
    this.getProjects();
  }

  getStory(): void {
    const id = Number(this.route.snapshot.paramMap.get('id'));
    this.storyService.getStory(id).subscribe((story) => (this.story = story));
  }

  getDevelopers(): void {
    this.developerService
      .getDevelopers()
      .subscribe((developers) => (this.developers = developers));
  }

  getProjects(): void {
    this.projectService
      .getProjects()
      .subscribe((projects) => (this.projects = projects));
  }

  goBack(): void {
    this.location.back();
  }

  save(): void {
    if (this.story) {
      this.storyService.updateStory(this.story).subscribe();
    }
  }
}
