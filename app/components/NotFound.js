import React, {Component} from 'react'
import html from '../content/notfound.md'

export default class NotFound extends Component {
  render() {
    return (
      <div class="NotFound container section content" dangerouslySetInnerHTML={{__html: html}}>
      </div>
    )
  }
}
