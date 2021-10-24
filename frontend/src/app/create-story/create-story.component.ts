import { Component, OnInit } from '@angular/core';
import { Developer } from '../developer';
import { DeveloperService } from '../developer.service';
import { StoryInput } from '../story';
import { StoryService } from '../story.service';

@Component({
  selector: 'app-create-story',
  templateUrl: './create-story.component.html',
  styleUrls: ['./create-story.component.css'],
})
export class CreateStoryComponent implements OnInit {
  developers: Developer[] = [];
  constructor(
    private storyService: StoryService,
    private developerService: DeveloperService
  ) {}

  ngOnInit(): void {
    this.getDevelopers();
  }

  addStory(
    title: string,
    description: string,
    time: number,
    developerName: string
  ): void {
    if (!title || !description || !time || !developerName) {
      return;
    }
    const storyInput: StoryInput = {
      title: title,
      description: description,
      time: time,
      developer_name: developerName,
    };
    this.storyService.addStory(storyInput).subscribe();
  }

  getDevelopers(): void {
    this.developerService
      .getDevelopers()
      .subscribe((developers) => (this.developers = developers));
  }
}
