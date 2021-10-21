import { Component, OnInit } from '@angular/core';
import { StoryInput } from '../story';
import { StoryService } from '../story.service';

@Component({
  selector: 'app-create-story',
  templateUrl: './create-story.component.html',
  styleUrls: ['./create-story.component.css'],
})
export class CreateStoryComponent implements OnInit {
  constructor(private storyService: StoryService) {}

  ngOnInit(): void {}

  addStory(title: string, description: string, time: number): void {
    if (!title || !description || !time) {
      return;
    }
    const storyInput: StoryInput = {
      title: title,
      description: description,
      time: time,
    };
    this.storyService.addStory(storyInput).subscribe();
  }
}
